import { defineStore } from 'pinia';
import axios from 'axios';

export class Equity {
  ticker: string;
  shares: number;
  equityPrice: number;
  priceLoaded: boolean;
  value: number;

  constructor(
    ticker: string,
    shares: number,
    equityPrice: number,
    priceLoaded: boolean,
    value: number
  ) {
    this.ticker = ticker;
    this.shares = shares;
    this.equityPrice = equityPrice;
    this.priceLoaded = priceLoaded;
    this.value = value;
  }
}

export class Property {
  type: string;
  id: number;
  description: string;
  price: number;
  lien: number;
  rate: number;
  years: number;
  monthsLeft: number;
  value: number;

  constructor(
    type: string,
    id: number,
    description: string,
    price: number,
    lien: number,
    rate: number,
    years: number,
    monthsLeft: number,
    value: number
  ) {
    this.type = type;
    this.id = id;
    this.description = description;
    this.price = price;
    this.lien = lien;
    this.rate = rate;
    this.years = years;
    this.monthsLeft = monthsLeft;
    this.value = value;
  }
}

export const usePortfolioStore = defineStore('portfolioStore', {
  state: () => {
    return {
      equities: [],
      realEstate: new Map<number, Property>(),
      // Equity table,
      equityRows: [] as Array<Equity>,

      equitiesTotal: 0,

      form: {
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
        RErows: [] as Array<Property>,

        // Property Form Data
        type: '',
        id: 0,
        description: '',
        price: 0,
        lien: 0,
        rate: 0.04,
        years: 30,
        monthsLeft: 360,
        value: 0,

        // Property Modify Form Data
        modtype: '',
        modid: 0,
        moddescription: '',
        modREprice: 0,
        modlien: 0,
        modrate: 0.04,
        modyears: 30,
        modmonthsLeft: 360,
        modvalue: 0,
      },
    };
  },
  actions: {
    addEquityApi() {
      axios
        .post(
          'http://localhost:8081/portfolioadd',
          JSON.stringify({
            ticker: this.form.ticker,
            shares: Number(this.form.shares),
            price: Number(this.form.equityPrice),
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentEquities();
        });
    },
    addPropertyAPI() {
      axios
        .post(
          'http://localhost:8081/propertyadd',
          JSON.stringify({
            type: this.form.type,
            id: Number(this.realEstate.size),
            description: this.form.description,
            price: Number(this.form.price),
            lien: Number(this.form.lien),
            rate: Number(this.form.rate),
            years: Number(this.form.years),
            monthsLeft: Number(this.form.monthsLeft),
            value: Number(this.form.value),
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentProperties();
        });
    },
    async importCurrentEquities() {
      try {
        this.equitiesTotal = 0;
        const res = await axios
          .get('http://localhost:8081/portfolio')
          .then((res) => {
            let tempticker = '';
            let tempshares = 0;
            let tempvalue = 0;
            let tempprice = 0;

            this.equities = [];

            for (let i = 0; i < res.data.length; i++) {
              tempticker = res.data[i].ticker;
              tempshares = res.data[i].shares;
              tempvalue = res.data[i].value;
              tempprice = res.data[i].price;
              this.equities.push(tempticker, tempshares, tempvalue, tempprice);
            }

            this.equityRows = [] as Array<Equity>;

            for (let i = 0; i < this.equities.length; i++) {
              if (typeof this.equities[i] == 'string') {
                this.equitiesTotal +=
                  this.equities[i + 2] * this.equities[i + 1];
                this.equityRows.push(
                  new Equity(
                    this.equities[i],
                    this.equities[i + 1],
                    this.equities[i + 3],
                    false,
                    this.equities[i + 2]
                  )
                );
              }
            }

            this.equityRows.shift();
          });
      } catch {
        console.log('Backend API is not available');
      }
    },
    async importCurrentProperties() {
      const res = await axios
        .get('http://localhost:8081/properties')
        .then((res: any) => {
          this.realEstate.set(
            res.Id,
            new Property(
              res.type,
              res.id,
              res.description,
              res.price,
              res.lien,
              res.rate,
              res.years,
              res.monthsLeft,
              res.value
            )
          );
          console.log(res);
          console.log(res.type);
        });
    },
    async modifyEquity() {
      console.log(Number(this.form.modcurrentprice));

      axios
        .post(
          'http://localhost:8081/portfoliomod',
          JSON.stringify({
            ticker: this.form.modcurrentticker,
            shares: Number(this.form.modcurrentshares),
            price: Number(this.form.modcurrentprice),
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentEquities();
        });
    },
    async deleteEquity() {
      axios
        .post(
          'http://localhost:8081/portfoliodel',
          JSON.stringify({
            ticker: this.form.delcurrentticker,
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentEquities();
        });
    },
  },
  getters: {
    grabEquities: (state) => {
      return [...state.equities];
    },
    grabRealEstate(state) {
      return [...state.realEstate];
    },
  },
});
