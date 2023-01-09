class Currency {
    constructor(private value: number) {
        if (value < 0) {
            throw new Error("0以上を指定してください")
        }
        this.value = value
    }
}

class Money {
    constructor(private amount: number, private currency: Currency) {
        if (amount < 0) {
            throw new Error("金額は0以上を指定してください")
        }
        this.amount = amount
        this.currency = currency
    }

    add(other: number): Money {
        const added = this.amount + other
        return new Money(added, this.currency)
    }
}

const currency = new Currency(3)
const money = new Money(3, currency)
const money2 = money.add(3)
console.log(money2)
console.log(money)