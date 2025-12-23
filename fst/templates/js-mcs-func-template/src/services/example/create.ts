

export async function createExampleFactory(model:string, ){

    function execute(params: {id:number, name:string}):Promise<string>{
        throw new Error("Implement")
    }
    
    return {execute}
    
}