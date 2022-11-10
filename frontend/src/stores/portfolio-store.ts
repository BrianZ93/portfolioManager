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
  ownership: number;
  date: string;
  tenants: Array<Tenant>;
  buildingexpenses: Array<Expense>;

  constructor(
    type: string,
    id: number,
    description: string,
    price: number,
    lien: number,
    rate: number,
    years: number,
    monthsLeft: number,
    value: number,
    ownership: number,
    date: string,
    tenants: Array<Tenant>,
    buildingexpenses: Array<Expense>
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
    this.ownership = ownership;
    this.date = date;
    this.tenants = tenants;
    this.buildingexpenses = buildingexpenses;
  }
}

export class PropertyRow {
  type: string;
  id: number;
  description: string;
  price: number;
  lien: number;
  rate: number;
  years: number;
  monthsLeft: number;
  value: number;
  equity: number;
  ownership: number;
  date: string;
  currentBalance: number;
  tenants: Array<Tenant>;
  buildingexpenses: Array<Expense>;

  constructor(
    type: string,
    id: number,
    description: string,
    price: number,
    lien: number,
    rate: number,
    years: number,
    monthsLeft: number,
    value: number,
    equity: number,
    ownership: number,
    date: string,
    currentBalance: number,
    tenants: Array<Tenant>,
    buildingexpenses: Array<Expense>
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
    this.equity = equity;
    this.ownership = ownership;
    this.date = date;
    this.currentBalance = currentBalance;
    this.tenants = tenants;
    this.buildingexpenses = buildingexpenses;
  }
}

export class Tenant {
  name: string;
  leasestart: string;
  leaseend: string;
  expenses: Array<Expense>;
  revenues: Array<Revenue>;
  unit: string;
  currenttenant: boolean;
  id: number;
  subid: number;

  constructor(
    name: string,
    leasestart: string,
    leaseend: string,
    expenses: Array<Expense>,
    revenues: Array<Revenue>,
    unit: string,
    currenttenant: boolean,
    id: number,
    subid: number
  ) {
    this.name = name;
    this.leasestart = leasestart;
    this.leaseend = leaseend;
    this.expenses = expenses;
    this.revenues = revenues;
    this.unit = unit;
    this.currenttenant = currenttenant;
    this.id = id;
    this.subid = subid;
  }
}

export class Expense {
  description: string;
  amount: number;
  date: string;

  constructor(description: string, amount: number, date: string) {
    this.description = description;
    this.amount = amount;
    this.date = date;
  }
}

export class Revenue {
  description: string;
  amount: number;
  date: string;

  constructor(description: string, amount: number, date: string) {
    this.description = description;
    this.amount = amount;
    this.date = date;
  }
}

export class Ticker {
  ticker: string;
  price: number;
  change: number;

  constructor(ticker: string, price: number, change: number) {
    this.ticker = ticker;
    this.price = price;
    this.change = change;
  }
}

export class Debt {
  id: number;
  title: string;
  type: string;
  amount: number;
  rate: number;
  term: number;
  payment: number;
  date: string;

  constructor(
    id: number,
    title: string,
    type: string,
    amount: number,
    rate: number,
    term: number,
    payment: number,
    date: string
  ) {
    this.id = id;
    this.title = title;
    this.type = type;
    this.amount = amount;
    this.rate = rate;
    this.term = term;
    this.payment = payment;
    this.date = date;
  }
}

export const usePortfolioStore = defineStore('portfolioStore', {
  state: () => {
    return {
      equities: [] as Array<Equity>,
      realEstate: new Map<number, Property>(),

      propertyRows: [] as Array<PropertyRow>,
      // Equity table,
      equityRows: [] as Array<Equity>,

      Debts: [] as Array<Debt>,

      // Chart Specific Data

      // Asset Totals
      equitiesTotal: 0,
      propertiesTotal: 0,
      debtsTotal: 0,
      // Property Type Totals
      primaryTotal: 0,
      secondaryTotal: 0,
      investmentTotal: 0,
      // Equities Type Totals
      equityTickers: [] as Array<string>,
      equityValues: [] as Array<number>,

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
        ownership: 0,
        REdate: '',

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
        modOwnership: 0,
        modREDate: '',
        modTenants: [] as Array<Tenant>,
        modExpenses: [] as Array<Expense>,

        // Property Delete Form Data
        delcurrentproperty: 0,

        // Selected Property
        selectedProperty: new Property(
          '',
          0,
          '',
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          '',
          [] as Array<Tenant>,
          [] as Array<Expense>
        ),

        newTenant: new Tenant(
          '',
          '',
          '',
          [] as Array<Expense>,
          [] as Array<Revenue>,
          '',
          false,
          0,
          0
        ),

        // Debt Form Data
        debtid: 0,
        debttitle: '',
        debttype: '',
        debtamount: 0,
        debtrate: 0,
        debtterm: 0,
        debtpayment: 0,
        debtdate: '',

        // Debt Modify Form Data
        debtmodid: 0,
        debtmodtitle: '',
        debtmodtype: '',
        debtmodamount: 0,
        debtmodrate: 0,
        debtmodterm: 0,
        debtmodpayment: 0,
        debtmoddate: '',

        // Debt Delete Form Data
        delcurrentdebt: 0,
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
      const rate = parseFloat(this.form.rate);
      const ownership = parseFloat(this.form.ownership);

      axios
        .post(
          'http://localhost:8081/propertyadd',
          JSON.stringify({
            type: this.form.type,
            id: Number(this.realEstate.size),
            description: this.form.description,
            price: Number(this.form.price),
            lien: Number(this.form.lien),
            rate: Number(rate),
            years: Number(this.form.years),
            monthsLeft: Number(this.form.monthsLeft),
            value: Number(this.form.value),
            ownership: Number(ownership),
            date: this.form.REdate,
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentProperties();
        });
    },
    addDebtAPI() {
      const rate = parseFloat(this.form.debtrate);
      axios
        .post(
          'http://localhost:8081/debtadd',
          JSON.stringify({
            title: this.form.debttitle,
            type: this.form.debttype,
            amount: Number(this.form.debtamount),
            rate: Number(rate),
            term: Number(this.form.debtterm),
            payment: Number(this.form.debtpayment),
            date: this.form.debtdate,
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentDebts();
        });
    },
    async importCurrentEquities() {
      try {
        const res = await axios
          .get('http://localhost:8081/portfolio')
          .then((res) => {
            this.equities = [] as Array<Equity>;

            for (let i = 0; i < res.data.length; i++) {
              this.equities.push(
                new Equity(
                  res.data[i].ticker as string,
                  res.data[i].shares as number,
                  res.data[i].price,
                  false,
                  res.data[i].value
                )
              );
            }

            this.equityRows = [] as Array<Equity>;

            this.equitiesTotal = 0;
            this.equityTickers = [] as Array<string>;
            this.equityValues = [] as Array<number>;
            for (const equity of this.equities) {
              this.equityRows.push(equity);
              this.equitiesTotal += equity.value * equity.shares;

              this.equityTickers.push(equity.ticker);
              this.equityValues.push(equity.value * equity.shares);
            }
          });
      } catch {
        console.log('Backend API is not available');
      }
    },
    async importCurrentProperties() {
      this.propertyRows = [] as Array<PropertyRow>;

      const res = await axios
        .get('http://localhost:8081/properties')
        .then((res: any) => {
          this.propertiesTotal = 0;

          for (let i = 0; i < res.data.length; i++) {
            console.log(res.data[i]);
            const tempTenants = [] as Array<Tenant>;
            if (res.data[i].tenants) {
              console.log(res.data[i].description);
              for (const tenant of Object.entries(res.data[i].tenants)) {
                const t = tenant[1] as Tenant;
                tempTenants.push(
                  new Tenant(
                    t.name,
                    t.leasestart,
                    t.leaseend,
                    t.expenses,
                    t.revenues,
                    t.unit,
                    t.currenttenant,
                    t.id,
                    t.subid
                  )
                );
              }
            }

            console.log(tempTenants);

            this.realEstate.set(
              res.data[i].id,
              new Property(
                res.data[i].type,
                res.data[i].id,
                res.data[i].description,
                res.data[i].price,
                res.data[i].lien,
                res.data[i].rate,
                res.data[i].years,
                res.data[i].monthsLeft,
                res.data[i].value,
                res.data[i].ownership,
                res.data[i].date,
                tempTenants,
                res.data[i].buildingexpenses
              )
            );

            // Creating a blank date object
            const dateObject = new Date();
            // Retrieving todays date
            const dd = String(dateObject.getDate()).padStart(2, 0);
            const mm = String(dateObject.getMonth() + 1).padStart(2, '0');
            const yyyy = dateObject.getFullYear();

            const todayString = (mm + '/' + dd + '/' + yyyy) as string;

            // Getting the current property lien to be amortized
            const amortizingProperty = this.realEstate.get(res.data[i].id);

            // Calculating how many payments to amortize
            const yearsElapsed =
              parseFloat(todayString.slice(6, 10)) -
              parseFloat(amortizingProperty?.date.slice(6, 10) as string);

            const monthsElapsed =
              parseFloat(todayString.slice(0, 2)) -
              parseFloat(amortizingProperty?.date.slice(0, 2) as string) +
              yearsElapsed * 12;

            const monthlyRate = res.data[i].rate / 100 / 12;
            const rate = res.data[i].rate / 100;
            const term = res.data[i].years * 12;
            const lien = res.data[i].lien;

            const factor = (1 + monthlyRate) ** term;
            const rateFactor = monthlyRate * factor;

            const termFactor = (1 + monthlyRate) ** term - 1;

            const payment = lien * (rateFactor / termFactor);

            let currentBalance = lien;
            for (let i = 1; i < monthsElapsed; i++) {
              const interest = (rate / 12) * currentBalance;

              currentBalance = currentBalance - (payment - interest);
            }

            this.propertyRows.push(
              new PropertyRow(
                res.data[i].type,
                res.data[i].id,
                res.data[i].description,
                res.data[i].price,
                res.data[i].lien,
                res.data[i].rate,
                res.data[i].years,
                res.data[i].monthsLeft,
                res.data[i].value,
                res.data[i].value - currentBalance,
                res.data[i].ownership,
                res.data[i].date,
                currentBalance,
                tempTenants,
                res.data[i].buildingexpenses
              )
            );

            this.propertiesTotal += res.data[i].value - res.data[i].lien;

            if (res.data[i].type == 'Primary Residence') {
              this.primaryTotal += res.data[i].value - res.data[i].lien;
            } else if (res.data[i].type == 'Second Home') {
              this.secondaryTotal += res.data[i].value - res.data[i].lien;
            } else if (res.data[i].type == 'Investment Property') {
              this.investmentTotal += res.data[i].value - res.data[i].lien;
            } else {
              console.log('property type not recognized');
            }
          }
        });
    },
    async importCurrentDebts() {
      try {
        const res = await axios
          .get('http://localhost:8081/debts')
          .then((res) => {
            this.Debts = [] as Array<Debt>;

            for (let i = 0; i < res.data.length; i++) {
              this.Debts.push(
                new Debt(
                  res.data[i].id as number,
                  res.data[i].title as string,
                  res.data[i].type as string,
                  res.data[i].amount as number,
                  res.data[i].rate as number,
                  res.data[i].term as number,
                  res.data[i].payment as number,
                  res.data[i].date as string
                )
              );
            }

            this.debtsTotal = 0;
            for (const debt of this.Debts) {
              this.debtsTotal += debt.amount;
            }
          });
      } catch {
        console.error('Backend API is not available');
      }
    },
    async modifyEquity() {
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
    async modifyProperty() {
      const rate = parseFloat(this.form.modrate);
      const ownership = parseFloat(this.form.modOwnership);

      axios
        .post(
          'http://localhost:8081/propertymod',
          JSON.stringify({
            id: Number(this.form.modid),
            description: this.form.moddescription,
            price: Number(this.form.modREprice),
            lien: Number(this.form.modlien),
            rate: Number(rate),
            years: Number(this.form.modyears),
            value: Number(this.form.modvalue),
            monthsleft: Number(this.form.modmonthsLeft),
            type: this.form.modtype,
            ownership: Number(ownership),
            date: this.form.modREDate,
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentProperties();
        });
    },
    async deleteProperty() {
      axios
        .post(
          'http://localhost:8081/propertydel',
          JSON.stringify({
            id: this.form.delcurrentproperty,
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentProperties();
        });
    },
    async modifyDebt() {
      const rate = parseFloat(this.form.debtmodrate);

      axios
        .post(
          'http://localhost:8081/debtmod',
          JSON.stringify({
            id: Number(this.form.debtmodid),
            title: this.form.debtmodtitle,
            type: this.form.debtmodtype,
            amount: Number(this.form.debtmodamount),
            rate: Number(rate),
            term: Number(this.form.debtmodterm),
            payment: Number(this.form.debtmodpayment),
            date: this.form.debtmoddate,
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentDebts();
        });
    },
    async deleteDebt() {
      axios
        .post(
          'http://localhost:8081/debtdel',
          JSON.stringify({
            id: this.form.delcurrentdebt,
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentDebts();
        });
    },
    async addTenant() {
      console.log(this.form.newTenant);
      console.log(this.form.selectedProperty.id);
      axios
        .post(
          'http://localhost:8081/tenantAdd',
          JSON.stringify({
            id: Number(this.form.selectedProperty.id),
            tenants: Array(this.form.newTenant) as Array<Tenant>,
          })
        )
        .then((response) => {
          console.log(response);
          this.importCurrentProperties();
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
    grabDebts: (state) => {
      return [...state.Debts];
    },
  },
});
