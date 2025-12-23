import type { Request, Response , NextFunction } from "express";
import type { Iexample } from "../../app/example.interfaces.js";



export class ExampleController {

    constructor(private example: Iexample){}

    create(req:Request, res:Response, nextFunction:NextFunction){

        const result = this.example.create(req.body)

        res.status(201).json(result)
    }
}
