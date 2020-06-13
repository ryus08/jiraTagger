# Investigation of IoC container libs:

## Evaluation

Things I'm taking into consideration:
* Whether code gen is needed
* If we have to manually construct everything, the container just is a place to put everything
* How to scope instances
* How much duplication of dependency definitions are needed, places such as
  * In the struct definition (of course)
  * Explicit annotations of injectability on the struct
  * Manually in a constructor 
  * Manually passing into the constructor
  * In a separate type definition configuration file
* Named instances (although just providing a couple different named interfaces on top of the type is probably the Go way to do it)
* Collection injection (get me all registered instances)

### Wire 

https://github.com/google/wire

Just define structs, constructors (called providers in Wire) and tell wire all the providers to use in an "injector". Code-gen driven to shake the dependency tree

### Dig

https://github.com/uber-go/dig

Not code generated, just create a container, tell it the providers you have, and go. Support for named injection, and collections. No scopes, but could set up manually with multiple instances, and has an MR open with all checks except approval and CLA signature ready.

### DI

https://github.com/sarulabs/di

Runtime and simple. Expects all registrations to directly interact with the container and be named. No dependency tree generation it would seem

### Container

https://github.com/golobby/container

Runtime, dependency tree shaking. No scopes and seems like a singleton container, so while they refer to "Request-scoped things" no idea how you could do it.

### Dingo

https://github.com/elliotchance/dingo

Code generation, config-based. Seems to generate the type definitions themselves, not just the constructors. Doesn't have "request scope" but has tutorial on how to use it to generate builders which would be used in the request scope


### IoC

https://github.com/gopub/ioc

Low usage, little docs, but seemingly full-featured without code generation. Not manual wiring, but needs annotations for dependency references

### Go IoC

https://github.com/shelakel/go-ioc

Low usage, little docs, but seemingly full-featured without code generation but with manual wiring


## Comparison

|           | Stars  | Scopes | Code Generated   | Manual wiring  | Config Based  | Named Instances | Collection injection |
|-----------|--------|--------|------------------|---|---|---|---|
|  Wire    | 3800    | ❌     | ✔️              | ❌   | ❌  | ❌| ❌
|  Dig    | 1400    |   ❌   |   ❌            | ❌  | ❌ |✔️ |✔️
| DI        | 300    | ✔️     | ❌              | ✔️  | ❌  || ❌
| Container | 100    | ❌     | ❌              | ❌  | ❌ || ❌
| Dingo     | 100    | ❌     | ✔️              |  ❌ | ✔️ || ❌
|  Ioc    | 6    |      |               |   |  || ❌
| Go-Ioc       | 15    |  ✔️    | ❌              | ✔️  |❌| ❌ | ❌