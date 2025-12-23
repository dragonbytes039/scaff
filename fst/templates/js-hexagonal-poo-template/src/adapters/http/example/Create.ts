import type { Request, Response, NextFunction } from "express";
import type { ICreateExampleUseCase } from "../../../internal/ports/useCases/example.interfaces.js";

export class CreateExampleController{

    constructor( private createExampleUseCase:ICreateExampleUseCase){}

    async execute (req:Request, res:Response, next:NextFunction){

        const data = {
            id:req.body.id,
            name:req.body.name,
            description:req.body.description
        }

        try {
            const result = this.createExampleUseCase.execute(data)

            res.status(200).json(result)

        } catch (error) {
            res.status(500).send("error")
        }

    }
    
}