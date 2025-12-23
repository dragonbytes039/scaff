import type { Request, Response , NextFunction } from "express";
import { createExampleFactory } from "../services/example/create.js";



export async function createExampleController(req:Request, res:Response, next:NextFunction){
        
        try {

            const result = (await createExampleFactory("insertModel o repo")).execute({id:1,name:"test"})

            res.status(200).send("create")
            
        } catch (error) {
            res.status(500).send("errror")
        }
        
    }