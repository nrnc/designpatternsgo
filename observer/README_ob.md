## Chapter 1: Observer pattern

### Publishers + Subscribers = Observer Pattern

### Definition
The Observer Pattern defines a one-to-many dependency between objects so that when one object changes state, all of its dependents are notified and updated automatically.

### Design Principles
- Strive for loosely coupled design between the objects that interact

Loosely coupled designs allow us to build flexible OO systems that can handle change because they minimize the interdependency between objects.


### Key Terms in Observer Pattern
1. **Publisher/Subject** : The Publisher issues events of interest to other objects. These events occur when the publisher changes its state or executes some behaviors. Publishers contain a subscription infrastructure that lets new subscribers join and current subscribers leave the list.
2. **Subscriber** : Subscribers perform some actions in response to notifications issued by the publisher. All of these classes must implement the same interface(subscriber interface) so the publisher isn’t coupled to concrete classes.
