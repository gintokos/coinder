class Coin {
    fullname = ''
    shortname = ''
    uniqid = 0
    likesAmount = 0
    dislikesAmount = 0
    photoUrl = ''
    network = ''
    currentPrice = 0
    volume24h = 0
    circulatingSupply = 0
    totalSupply = 0
    maxSupply = 0
    contract = ''
    burses = []
    percentChange = {
        hour: 0,
        day: 0,
        week: 0,
        month: 0,
        year: 0
    }

    constructor(data) {
        this.update(data)
    }

    update(data) {
        if (!data) return

        if (data.percentChange) {
            this.percentChange = { ...this.percentChange, ...data.percentChange }
        }

        if (data.burses) {
            this.burses = [...data.burses]
        }

        Object.assign(this, data)
    }
    
    getuniqname() {
        if (this.network === undefined) {
            return `${this.fullname}(${this.shortname})`.trim()
        }
        return `${this.fullname}(${this.network})`.trim()
    }
}

export default Coin