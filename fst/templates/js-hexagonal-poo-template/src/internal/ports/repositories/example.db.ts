import type { Example } from "../../domain/entities/Example.js";


export interface IExampleDbRepository {
    
    create:(example:Example)=>Promise<Example>, 
    deleteById:(example:Example['id'])=>Promise<boolean>,

}