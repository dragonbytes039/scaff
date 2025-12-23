import type { Example } from "../domain/entities.js";
import type { DbRepository } from "../repositories/db.interfaces.js";
import type { Iexample } from "./example.interfaces.js";


export class ExampleUseCase  implements Iexample{

    constructor(private repo :DbRepository){}
    create(example: Example): Example {
        throw new Error("Method not implemented.");
    }
    
}