# GoDungeon

## Architecture Notes

### Game Layers

**Game** contains data about the game world. The data itself is not smart
enough to do anything asides return some basic information about itself.

**Controllers** contains objects that directly manipulate game data. The
objects are stateless. They make decisions about how actions effect game
objects, but don't make any decisions about when to execute those changes.

**Updaters** contains objects that are polled once each refresh. They make
decisions about what should be happening to objects based on the current
game state, and pass those decisions onto the appropriate controller.

**Renderers** render an object to a canvas layer. They are called once per
refresh but should only change the layer state if the object/s they render
have changed in some way.

**Scenes** coordinate which updaters and renderers should be running at
any given time.

## Support Packages

**io** contains objects that load dungeon data, and setup a game
environment based on that data. Maps should be created using the mapmaker
library.

**input** contains representations of user input, and provides the
interface for retrieving input events.

