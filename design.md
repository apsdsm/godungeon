dungeon game

map:
- import a map and display it on a canvas layer
- update the map
- position entities on the map
- move around the map and let the entities see each other and respond to each other.

IO Domain
- load the map file
- load any other required files
- return the json unmarshaled in the formats supplied by other domains

Map Domain
- take an unmarshaled map and display it on a layer

Input Domain
- take keyboard input and emit events when something happens

Entities Domain
- load the player and position on map
- load the entities
    - create using prototypes and instances
    - assign controllers
    - postion entities on map
- keep a list of the entities
- send entities the update message
- entity controllers in here

Pathfinding
- allow entities to find accessible tiles in relation to themselves
- allow entities to get a range of tiles representing how far they can see

entity <- each turn
- needs to figure out what it wants to do
- what can I see?
- what is my current priority? <- if I see a target do I chase it? ru for help?
- am I close enough to attack?

should make a controller class for each enemy. The entity itself is dumb, but by attaching a different AI I can control how it works. More specifically the AI should be set by parameters in the entity file.

mob
    player <- playerController
    enemy <- enemyController

# Play Screen Flow

All entities in the game are assigned a controller. The controller is a class that is responsible for making decisions for the enemy, based on the enemy's current state. There's no need for multiple controllers to be assigned to enemies of the same class - the same controller can work for all of them. So in principle, a player controller, and an enemy controller. There may later be an NPC controller.

Inside of that, each entitiy is assigned an update order, which is just their arrangement in an array. The player is always at the head of this array. The order of the other monsters will be fine as a random allocation for now. The idea is that all entities (player included) move at the same time, so there's no real reason to worry about initiative asides who attacks first (movement is not included in initiative).

when the player presses a key, they move in a direction. At that time other entities will also move.
Only entities that are "active" will be updated. When an entity is at rest it is no longer active. most enemies will remain inactive unless the player gets too close to them. this distance is decided by some magic number somewhere.

after there is a player input, the rest of the game moves. If this isn't fast enough it's going to make a mess of everything, and input is going to feel slow.

1. player input phase
    - movement mode
        - press a movement key
            - if movement is possible in that direction
                - move one space in that direction
            - if there is an enemy in that direction
                - melee attack in that direction
            - if there is a door in that direction
                - if the player has the key
                    - open the door
                    - display the OPEN_DOOR message
                - if the player doesn't have the key
                    - display the OPEN_DOOR_FAILED message
            - if there is a friendly in that direction
                - talk to the friendly
        - press the ranged attack key (R, or maybe whatever the assign weapon slot is e.g., 1)
            - switch to range attack mode (highlight nearest hittable oppenent)
        - press the examine mode key (E)
            - go to examine mode
        - press the inventory key (I)
            - go the equipment screen (can swap out equipment and check items, health, etc)
    - ranged weapon attack mode
        - press the LEFT or RIGHT key (A, D)
            - cycle to the prev or next hittable opponent
        - press the CANCEL key (C)
          - return to movement mode
        - press the FIRE key (F)
            - initiate a ranged attack on the target
    - examine mode (E)
        - press the DIRECTION keys (W, A, S, D) to move the focus cursor inside of visible range
            - if the cursor is over something with a description
                - display the desscription

The player needs to be able to:
- move
- meelee attack
- switch weapons
- use keys
- talk to people

The game needs to:
- [x] load a map
- [ ] display the map part of the dungeon
    - [ ] basic display - put the runes on the scree
    - [ ] show entities
    - [ ] center map
    - [ ] scroll map
- [ ] put the player on the map in the starting position
- [ ] put enemies on the map in their starting position

The enemy controllers*
- needs to move the enemies and make choices for them.



 
