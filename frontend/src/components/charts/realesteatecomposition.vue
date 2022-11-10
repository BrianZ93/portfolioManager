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

<script>
import { usePortfolioStore } from 'src/stores/portfolio-store';

const portfolioStore = usePortfolioStore();

export default {
  name: 'realEstateChart',
  data() {
    return {
      series: [
        portfolioStore.primaryTotal,
        portfolioStore.secondaryTotal,
        portfolioStore.investmentTotal,
      ],

      chartOptions: {
        tooltip: {
          y: {
            formatter: function (value) {
              return value.toLocaleString('en-US', {
                style: 'currency',
                currency: 'USD',
              });
            },
          },
        },
        labels: [
          'Primary Residence',
          'Second Home(s)',
          'Investment Properties',
        ],

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
                show: true,
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
