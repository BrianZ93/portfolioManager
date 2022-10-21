import { defineStore } from 'pinia';
import axios from 'axios';

export const useIndexStore = defineStore('indexStore', {
  state: () => {
    return {
      indexes: [0,0,0,0,0,0,0,0],
      ticker: '',
      shares: 0
    }
  },
  actions: {
    async getData () {
      const res = await axios.get('http://localhost:8081/index')
      .then((res) => {
        this.indexes[0] = JSON.parse(res.data[0].price)
        this.indexes[1] = JSON.parse(res.data[0].change)
        this.indexes[2] = JSON.parse(res.data[1].price)
        this.indexes[3] = JSON.parse(res.data[1].change)
        this.indexes[4] = JSON.parse(res.data[2].price)
        this.indexes[5] = JSON.parse(res.data[2].change)
        this.indexes[6] = JSON.parse(res.data[3].price)
        this.indexes[7] = JSON.parse(res.data[3].change)
        for (let i=0;i < 4; i++) {
          console.log(res.data[i].price)
          console.log(res.data[i].change)
        }
      })
    }
  },
  getters: {
    data: (state) => [...state.indexes],
  }
});

