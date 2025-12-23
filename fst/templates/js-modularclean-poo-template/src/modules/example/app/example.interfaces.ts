import type { Example } from "../domain/entities.js";

export interface Iexample {
    create (data: Example):Example
}