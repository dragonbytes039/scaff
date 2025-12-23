import type {Request, Response, nextFunction} from "express"
import { MongoRepo } from "./repositories/example.mongo.repo.js";
import { ExampleController } from "./adapters/express/Example.controllers.js";
import { ExampleUseCase } from "./app/example.useCase.js";



const examplemongoRepo = new MongoRepo()
const exampleUseCase = new ExampleUseCase(examplemongoRepo)
export const exampleControllerBuild = new ExampleController(exampleUseCase)

    
