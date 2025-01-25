package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area/?offset=40&limit=20",
			val: []byte(`solaceon-ruins-b3f-d solaceon-ruins-b3f-e
solaceon-ruins-b4f-a
solaceon-ruins-b4f-b
solaceon-ruins-b4f-c
solaceon-ruins-b4f-d
solaceon-ruins-b5f
sinnoh-victory-road-1f
sinnoh-victory-road-2f
sinnoh-victory-road-b1f
sinnoh-victory-road-inside-b1f
sinnoh-victory-road-inside
sinnoh-victory-road-inside-exit
ravaged-path-area
oreburgh-gate-1f
oreburgh-gate-b1f
stark-mountain-area
stark-mountain-entrance
stark-mountain-inside
sendoff-spring-area`),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/?offset=60&limit=20",
			val: []byte(`turnback-cave-pillar-1
turnback-cave-pillar-2
turnback-cave-pillar-3
turnback-cave-before-pillar-1
turnback-cave-between-pillars-1-and-2
turnback-cave-between-pillars-2-and-3
turnback-cave-after-pillar-3
snowpoint-temple-1f
snowpoint-temple-b1f
snowpoint-temple-b2f
snowpoint-temple-b3f
snowpoint-temple-b4f
snowpoint-temple-b5f
wayward-cave-1f
wayward-cave-b1f
ruin-maniac-cave-0-9-different-unown-caught
ruin-maniac-cave-10-25-different-unown-caught
maniac-tunnel-26-plus-different-unown-caught
trophy-garden-area
iron-island-area`),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area/?offset=60&limit=20", []byte(`turnback-cave-pillar-1
turnback-cave-pillar-2
turnback-cave-pillar-3
turnback-cave-before-pillar-1
turnback-cave-between-pillars-1-and-2
turnback-cave-between-pillars-2-and-3
turnback-cave-after-pillar-3
snowpoint-temple-1f
snowpoint-temple-b1f
snowpoint-temple-b2f
snowpoint-temple-b3f
snowpoint-temple-b4f
snowpoint-temple-b5f
wayward-cave-1f
wayward-cave-b1f
ruin-maniac-cave-0-9-different-unown-caught
ruin-maniac-cave-10-25-different-unown-caught
maniac-tunnel-26-plus-different-unown-caught
trophy-garden-area
iron-island-area`))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area/?offset=60&limit=20")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area/?offset=60&limit=20")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}