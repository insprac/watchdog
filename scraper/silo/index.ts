import puppeteer from "puppeteer";

interface RawSiloToken {
  name: string;
  totalDeposited: string;
  availableToBorrow: string;
  utilization: string;
}

export interface SiloToken {
  name: string;
  totalDeposited: number;
  availableToBorrow: number;
  utilization: number;
}

export async function scrapeSiloTokens(address: string): Promise<SiloToken[]> {
  const browser = await puppeteer.launch({
    headless: "new",
  });
  const page = await browser.newPage();

  try {
    await page.goto("https://app.silo.finance/silo/" + address, {
      waitUntil: "networkidle0",
    });

    const siloTokens: RawSiloToken[] = await page.evaluate(() => {
      const tokenCards = document.querySelectorAll('[data-cy="asset-card"]');
      const tokens: RawSiloToken[] = [];

      tokenCards.forEach((card) => {
        const name =
          card.querySelector('[data-cy="asset-name"]')?.textContent || "";
        const totalDeposited =
          card.querySelector('[data-cy="total-deposited"] [data-cy="top"]')
            ?.textContent || "";
        const availableToBorrow =
          card.querySelector('[data-cy="left-to-borrow"] [data-cy="top"]')
            ?.textContent || "";
        const utilization =
          card.querySelector('[data-cy="utilization"]')?.textContent || "";

        tokens.push({
          name,
          totalDeposited,
          availableToBorrow,
          utilization,
        });
      });

      return tokens;
    });

    await browser.close();

    return siloTokens.map((token) => ({
      name: token.name,
      totalDeposited: parseSize(token.totalDeposited),
      availableToBorrow: parseSize(token.availableToBorrow),
      utilization: parseFloat(token.utilization),
    }));
  } catch (error) {
    await browser.close();
    throw error;
  }
}

function parseSize(size: string): number {
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
