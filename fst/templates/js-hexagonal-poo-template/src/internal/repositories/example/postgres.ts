import type { Example } from "../../domain/entities/Example.js"
import type { IExampleDbRepository } from "../../ports/repositories/example.db.js"

export class ExamplePostgresRepository implements IExampleDbRepository{

    async create(example:Example){

        return example
    }
    async deleteById(example:Example['id']){

        return false
    }

} 