## Chapter 1: Strategy pattern

### Definiton:

The Strategy Pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

### Design principles
- Identify the aspects of application that changes, separate them from what stays the same
- program to an interface, not to an implementation
- Has-a can be better than Is-A
- Favour composition over inheritence

### Key terms in Strategy pattern

1. Context: The context maintains a reference to one of the concrete strategy(Algorithm) and it calls the concrete objects implementaion through the interface that the concret object implements. In our example **Duck** class is the context
2. Strategy : Strategy is an interface which is common to all the concrete implementation. It declares a method which is implemented by all the concrete strategies. This method is used by the context to execute the strategy. FlyBehaviour and QuackBehaviour interfaces are the Strategy interface
3. Concrete Strategy : The concrete strategy implements the strategy interface. FlyWithWings, Squeak are the concrete strategies
4. Client: It creates the specific concrete strategy and passes to the context. The context exposes a setter which lets client replace the strategy at the run time.s