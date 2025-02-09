class Coin {
   constructor(data) {
       this.id = 0
       this.name = ''
       this.symbol = ''
       this.slug = ''
       this.likesAmount = 0
       
       this.lastUpdated = null
       this.dateAdded = null
       this.dateLaunched = null

       this.price = 0
       this.volume24h = 0
       this.volumeChange24h = 0
       this.marketCap = 0
       this.marketCapDominance = 0
       this.fullyDilutedMarketCap = 0

       this.percentChange = {
           hour: 0,
           day: 0,
           week: 0
       }

       this.urls = {
           website: '',
           technicalDoc: '',
           twitter: '',
           reddit: '',
           messageBoard: '',
           announcement: '',
           chat: '', 
           explorer: '',
           sourceCode: ''
       }

       this.logo = ''
       this.description = ''

       this.update(data)
   }

   update(data) {
       if (!data) return
       Object.assign(this, data)

       return this
   }
}

export default Coin