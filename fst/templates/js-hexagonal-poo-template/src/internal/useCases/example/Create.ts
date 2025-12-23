import type { IExampleDbRepository } from "../../ports/repositories/example.db.js";
import type { CreateExampleUseCaseDto } from "../../ports/useCases/example.dtos.js";
import type { ICreateExampleUseCase } from "../../ports/useCases/example.interfaces.js";


export class CreateExampleUseCase implements ICreateExampleUseCase {

    constructor(private exampleRepository: IExampleDbRepository ){

    }

    async execute(data: CreateExampleUseCaseDto) {
        return await this.exampleRepository.create(data)
    }

}