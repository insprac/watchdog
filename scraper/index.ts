import { scrapeSiloTokens } from "./silo";

async function main() {
  const siloAddress = process.argv[2];
  if (!siloAddress) {
    throw new Error(
      `A silo address argument is required\nUsage: node ./index.js <silo_address>`
    );
  }
  const siloTokens = await scrapeSiloTokens(siloAddress);
  console.log(formatJson(siloTokens));
}

function formatJson(obj: object): string {
  return JSON.stringify(obj, undefined, 2);
}

main().catch(console.error);
