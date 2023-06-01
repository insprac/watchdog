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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.scrapeSiloTokens = void 0;
const puppeteer_1 = __importDefault(require("puppeteer"));
function scrapeSiloTokens(address) {
    return __awaiter(this, void 0, void 0, function* () {
        const browser = yield puppeteer_1.default.launch({
            headless: "new",
        });
        const page = yield browser.newPage();
        try {
            yield page.goto("https://app.silo.finance/silo/" + address, {
                waitUntil: "networkidle0",
            });
            const siloTokens = yield page.evaluate(() => {
                const tokenCards = document.querySelectorAll('[data-cy="asset-card"]');
                const tokens = [];
                tokenCards.forEach((card) => {
                    var _a, _b, _c, _d;
                    const name = ((_a = card.querySelector('[data-cy="asset-name"]')) === null || _a === void 0 ? void 0 : _a.textContent) || "";
                    const totalDeposited = ((_b = card.querySelector('[data-cy="total-deposited"] [data-cy="top"]')) === null || _b === void 0 ? void 0 : _b.textContent) || "";
                    const availableToBorrow = ((_c = card.querySelector('[data-cy="left-to-borrow"] [data-cy="top"]')) === null || _c === void 0 ? void 0 : _c.textContent) || "";
                    const utilization = ((_d = card.querySelector('[data-cy="utilization"]')) === null || _d === void 0 ? void 0 : _d.textContent) || "";
                    tokens.push({
                        name,
                        totalDeposited,
                        availableToBorrow,
                        utilization,
                    });
                });
                return tokens;
            });
            yield browser.close();
            return siloTokens.map((token) => ({
                name: token.name,
                totalDeposited: parseSize(token.totalDeposited),
                availableToBorrow: parseSize(token.availableToBorrow),
                utilization: parseFloat(token.utilization),
            }));
        }
        catch (error) {
            yield browser.close();
            throw error;
        }
    });
}
exports.scrapeSiloTokens = scrapeSiloTokens;
function parseSize(size) {
    let scale = 1;
    if (typeof size === "string") {
        switch (size.slice(-1).toUpperCase()) {
            case "K":
                scale = 1000;
                break;
            case "M":
                scale = 1000000;
                break;
            case "B":
                scale = 1000000000;
                break;
        }
        return parseFloat(size) * scale;
    }
    return parseFloat(size);
}
