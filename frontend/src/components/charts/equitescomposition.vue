<template>
  <div class="chart-wrap">
    <div id="chart">
      <apexchart
        type="donut"
        width="480"
        :options="chartOptions"
        :series="series"
      ></apexchart>
    </div>
  </div>
</template>

<script lang="ts">
import { usePortfolioStore } from 'src/stores/portfolio-store';

const portfolioStore = usePortfolioStore();

export default {
  name: 'equitiesChart',
  data() {
    return {
      series: [...portfolioStore.equityValues],

      chartOptions: {
        tooltip: {
          y: {
            formatter: function (value: number) {
              return value.toLocaleString('en-US', {
                style: 'currency',
                currency: 'USD',
              });
            },
          },
        },
        labels: [...portfolioStore.equityTickers],
        chart: {
          width: 380,
          type: 'donut',
        },
        dataLabels: {
          enabled: true,
          style: {
            fontSize: '14px',
            fontFamily: 'Helvetica, Arial, sans-serif',
            fontWeight: 'bold',
            colors: ['#FFFFFF'],
          },
        },
        responsive: [
          {
            breakpoint: 480,
            options: {
              chart: {
                width: 200,
              },
              legend: {
                show: false,
              },
            },
          },
        ],
        legend: {
          title: 'assets',
          onItemHover: {
            highlightDataSeries: true,
          },
          position: 'bottom',
          offsetY: 50,
          height: 0,
        },
      },
    };
  },
};
</script>
