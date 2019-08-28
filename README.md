**Laeron** is the MUD game that I, personally, have wanted to build for many
years. I've constantly started projects that veered towards open-ended
features/capabilities like [DragonMUD](github.com/bbuck/dragon-mud); however,
the time commitment involved in solving the generic use case can be more
significant, especially when it has to be completed before development on the
game can actually begin.

So I have made the decision to build an opinionated MUD server that fulfills
the needs/desires and goals of **Laeron** directly so that I can build the game
alongside the development of the engine and let the game's needs drive the
development of this server. As with most of my tools that I end up building,
the server itself will be available to everyone so that they can build their
own experiences on top of the same rulesets as **Laeron** but with their own
flavor or hack their own unique behaviors in to add that extra personal touch.

Projects like DragonMUD will eventually be completed and be the entirely
open-ended server that let's potential server developers pick and choose and
completely customize with ease.

### Notable Features

This is (and for a while, will be) an incomplete lets of features/behaviors that
this server will exhibit in it's default state. It will also kind of serve as
a rough picture of a roadmap.

**NOTE:** My knowledge of the internals of typical and classic MUD and MU\*
variant projects is extremely limited as I've always aspired to build my own vs.
learning/building off of another engine. So forgive any incorrect commentary
towards "classic" behavior and feel free to open an issue describing the where
the assumption is invalid so these blurbs can be updated to be more accurate
and specific.

- [ ] **Classic Rooms**, basic 10 directional traversal (N, W, S, E, NW, NE, SW,
      SE, Up, Down)
  - [ ] **Instanced Rooms**, specialized rooms that, instead of behaving in a
        "classic" manner an intanced rooms spawns an "Instance" of itself when a
        player enters and associates a party (which could just be a single player)
        to the instance before eventually killing itself.
  - [ ] **Node Based Room Growth**, rooms can associated to "nodes" which can
        map to an area (named group of rooms) or multiple areas, etc... which
        means that certain activities inside of node connected room will affect
        the node associated to it. (e.g. mining resources from a forest will
        reduce a total "resource count."
  - [ ] **Phased Rooms**, each room can have several phases that are switched
        on or off by various values (maybe player accomplishments, global
        changes, etc...) to enable the feeling of growth and effect that players
        have on the world they inhabit (e.g. clearing a village of enemies and
        eventually getting NPCS, or clearing an area of once dense forest into
        a field of stumps)
- (many, many more to come)
