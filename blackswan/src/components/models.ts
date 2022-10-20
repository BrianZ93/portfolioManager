export class PropertySummary {
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
