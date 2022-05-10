# Clean Architecture Skeleton

This repository is skeleton to work to clean-architecture in Go Lang.  
Based on `DDD, Interfaces and Adapters`. Example its a cache and database repository  
Example with `Unit tests`. No external dependencies(Database, cache, api, etc)  
Development based on `TDD`  


- `CMD`: Start App(Rest, Grcp and Subscribers)  
- `Domain`: Core domain and validators business rules. Factorys and interfaces repositories
- `Infra`: Repository implementation. Database, Cache and Messaging connection(Adapters)
- `Modules`: Wirrings for modules. Dependency injection
- `Usecase`: functions, dto(Presenters) 

