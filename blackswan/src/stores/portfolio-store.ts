import { defineStore } from 'pinia';
import axios from 'axios';

export class Equity {

  ticker: string
  shares: number
  equityPrice: number
  priceLoaded: boolean
  value: number

  constructor(ticker: string, shares: number, equityPrice: number, priceLoaded: boolean, value: number) {
    this.ticker = ticker;
    this.shares = shares;
    this.equityPrice = equityPrice;
    this.priceLoaded = priceLoaded;
    this.value = value;
  }
}

export const usePortfolioStore = defineStore('portfolioStore', {
  state: () => {
    return {
      equities: [],
      realEstate: [],
      // Equity table,
      equityRows: [{
        ticker: '',
        shares: 0,
        equityPrice: 0,
        priceLoaded: false,
        value: 0,
      }],

      form: ({
        // Equity Form Data
        ticker: '',
        shares: 0,
        equityPrice: 0,

        // Equity Modify Form Data
        modticker: '',
        modshares: 0,
        modprice: 0,
        modcurrentticker: '',
        modcurrentshares: 0,
        modcurrentprice: 0,

        // Equity Delete Form Data
        delcurrentticker: '',

        // Real Estate Form Data
        description: '',
        price: 0,
        lien: 0,
        rate: .04,
        years: 30,
        monthsLeft: 360,
        value: 0
      })
    }
  },
  actions: {
    addEquityApi() {
        axios.post('http://localhost:8081/portfolioadd',
          JSON.stringify({
            ticker: this.form.ticker,
            shares: Number(this.form.shares),
            price: Number(this.form.equityPrice),
          }))
          .then(response => {
            console.log(response)
            this.importCurrentEquities()
          })
    },
    async importCurrentEquities() {
      try {
        const res = await axios.get('http://localhost:8081/portfolio')
        .then(res => {

          let tempticker = ''
          let tempshares = 0
          let tempvalue = 0
          let tempprice = 0

          this.equities = []

          for (let i = 0; i < res.data.length; i++) {
            tempticker = res.data[i].ticker
            tempshares = res.data[i].shares
            tempvalue = res.data[i].value
            tempprice = res.data[i].price
            this.equities.push(tempticker, tempshares, tempvalue, tempprice)
          }

          console.log('equities:', this.equities)

          this.equityRows = [{
            ticker: '',
            shares: 0,
            equityPrice: 0,
            priceLoaded: false,
            value: 0
          }]

          for (let i = 0; i < this.equities.length; i++) {
            if (typeof this.equities[i] == 'string') {
              this.equityRows.push(new Equity(this.equities[i], this.equities[i+1], this.equities[i+3], false, this.equities[i+2]))
            }
          }

          this.equityRows.shift()

          console.log(res)

        })
      } catch {
        (console.log('Backend API is not available'))
      }
    },
    async modifyEquity() {

      console.log(Number(this.form.modcurrentprice))

      axios.post('http://localhost:8081/portfoliomod',
      JSON.stringify({
        ticker: this.form.modcurrentticker,
        shares: Number(this.form.modcurrentshares),
        price: Number(this.form.modcurrentprice),
      }))
      .then(response => {
        console.log(response)
        this.importCurrentEquities()
      })

    },
    async deleteEquity() {
      axios.post('http://localhost:8081/portfoliodel',
      JSON.stringify({
        ticker: this.form.delcurrentticker
      }))
      .then(response => {
        console.log(response)
        this.importCurrentEquities()
      })
    }
  },
  getters: {
    grabEquities: (state) => {
      return [...state.equities]
    },
    grabRealEstate (state) {
      return [...state.realEstate]
    }
  },
});
