Scaley is a custom autoscaling solution for Engine Yard Cloud environments.

## Glossary ##

* ***Group***: A representation of an autoscaling group; made up of Permanent Servers, Scaling Servers, a Scaling Script, and a Strategy
* ***Individual***: A Stragegy that dictates that only a single Scaling Server is to be scaled during a given Scaling Event
* ***Legion***: A Strategy that dictates that all Scaling Servers in the Group are to be scaled in the same direction during a given Scaling Event
* ***Permanent Server***: A server that is considered part of the group, but is never affected by a Scaling Event
* ***Scaling Server***: A server that is part of the group and is a candidate for state changes during a Scaling Event
* ***Scaling Script***: An external script that determines if the Group should be scaled up or down at any given time, reflected by its return code (1 = down, 2 = up, all else = no change)
* ***Scaling Event***: An attempt to scale the Group up or down, locking the group for all other operations until complete
* ***Strategy***: The manner in which the Group is to be scaled (known: Individual, Legion; default: Legion)

## Methodology ##

Scaley is meant to be run periodically for each configured group. During each run, Scaley executes the Scaling Script associated with the Group and does the following based on the result of that run:

    * If an upscale is desired, attempt to scale the group up. Log any scaling errors as critical errors, and log a lack of capacity in the group as a warning.
    * If a downscale is desired, attempt to scale the group down. Log any scaling errors as critical errors.
    * If no change is desired, do not attempt to scale the group.

# History #

* v0.1.0 - Initial release
