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
4. Client: It creates the specific concrete strategy and passes to the context. The context exposes a setter which lets client replace the strategy at the run time.

### Applcability
- Use the strategy pattern when you want to use different variants of an algorithm with in an object and be able to switch from one algortihm to another at runtime.
  - Eg. Suppose let's consider n x n tic tac toe. You can decide winner based on different conditions. Each condition you can treat as different strategy(algorithm). You can decide the winning strategy at run time.


### How to implement

1. In the context class, identify an algorithm that’s prone to frequent changes. It may also be a massive conditional that selects and executes a variant of the same algorithm at runtime.
2. Declare the strategy interface common to all variants of the algorithm.
3. One by one, extract all algorithms into their own classes. They should all implement the strategy interface.
4. In the context class, add a field for storing a reference to a strategy object. Provide a setter for replacing values of that field. The context should work with the strategy object only via the strategy interface. The context may define an interface which lets the strategy access its data.
5. Clients of the context must associate it with a suitable strategy that matches the way they expect the context to perform its primary job.


### References:

1. https://refactoring.guru/design-patterns/strategy
2. Head first design patterns 2nd ed.

