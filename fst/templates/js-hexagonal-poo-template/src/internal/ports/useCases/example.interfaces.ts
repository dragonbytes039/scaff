
import type { Example } from "../../domain/entities/Example.js"
import type { CreateExampleUseCaseDto, DeleteExampleUseCaseDto } from "./example.dtos.js"


export interface ICreateExampleUseCase {

    execute:(data: CreateExampleUseCaseDto)=>Promise<Example>

}
export interface IdeleteExampleUseCase {

    execute:(data: DeleteExampleUseCaseDto)=>Promise<boolean>

}