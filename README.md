Scaley is a custom autoscaling solution for Engine Yard Cloud environments.

## Glossary ##

* ***Group***: A representation of an autoscaling group; made up of Permanent Servers, Scaling Servers, a Scaling Script, and a Strategy
* ***Individual***: A Stragegy that dictates that only a single Scaling Server is to be scaled during a given Scaling Event
* ***Legion***: A Strategy that dictates that all Scaling Servers in the Group are to be scaled in the same direction during a given Scaling Event
* ***Permanent Server***: A server that is considered part of the group, but is never affected by a Scaling Event
* ***Scaling Server***: A server that is part of the group and is a candidate for state changes during a Scaling Event
* ***Scaling Script***: An external script that determines if the Group should be scaled up or down at any given time, reflected by its return code (0 = up, !0 = down)
* ***Scaling Event***: An attempt to scale the Group up or down, locking the group for all other operations until complete
* ***Strategy***: The manner in which the Group is to be scaled (known: Individual, Legion; default: Legion)

## Methodology ##

Scaley is meant to be run periodically for each configured group. During each run, the following checks happen:

1. What is the current state of the Group?
2. In which direction does the Scaling Script say that we should scale if a Scaling Event happens?
3. What was the Scaling Script result in the last Scaley run?
4. Compare those two results
    * If they are the same and the group can scale in that direction, start a Scaling Event to scale the group in the desired direction.
    * If they are the same, but the group cannot scale in that direction, record the most recent Scaling Script result, log the issue, and exit.
    * If they are different, record the most recent Scaling Script result and exit.

# History #

* v0.1.0 - Initial release
