const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});

const currencies = {
    USD: 1.0,
    JPY: 113.5,
    EUR: 0.89,
    RUB: 74.36,
    GBP: 0.75
};

function printGreeting() {
    console.log("Welcome to Currency Converter!");
}

function printRates() {
    for(const curr in currencies) {
        console.log(`1 USD equals  ${currencies[curr]} ${curr}`);
    }
}

function printCurrencies() {
    console.log("I can convert USD to these currencies: "
        + Object.keys(currencies).join(", "));
}

function actionConvert() {
    readline.question(
        "Type the currency you wish to convert: ", currencyFrom => {
            const currFromValue = currencies[currencyFrom];

            if(currFromValue === undefined) {
                console.log("Unknown currency");
                readline.close();
                return;
            }

            readline.question("To: ", currencyTo => {
                const currToValue = currencies[currencyTo];

                if(currToValue === undefined) {
                    console.log("Unknown currency");
                    readline.close();
                    return;
                }

                readline.question("Amount: ", amount => {
                    readline.close();

                    amount = parseFloat(amount);

                    if(isNaN(amount)) {
                        console.log("The amount has to be a number");
                        return;
                    }

                    if(amount < 1) {
                        console.log("The amount cannot be less than 1");
                        return;
                    }

                    let result = (currToValue / currFromValue * amount).toFixed(4);
                    console.log(`Result: ${amount} ${currencyFrom} equals ${result} ${currencyTo}`);
                });
            });
    });
}

function main() {
    printGreeting();
    printRates();
    printCurrencies();
    actionConvert();
}

main();
