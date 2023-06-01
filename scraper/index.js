"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const silo_1 = require("./silo");
function main() {
    return __awaiter(this, void 0, void 0, function* () {
        const siloAddress = process.argv[2];
        if (!siloAddress) {
            throw new Error(`A silo address argument is required\nUsage: node ./index.js <silo_address>`);
        }
        const siloTokens = yield (0, silo_1.scrapeSiloTokens)(siloAddress);
        console.log(formatJson(siloTokens));
    });
}
function formatJson(obj) {
    return JSON.stringify(obj, undefined, 2);
}
main().catch(console.error);
