<template>
  <q-page>
    <h1>{{ title }}</h1>

    <div>
      <div class="q-pa-md">
        <q-table
          title="Equities"
          :columns="columns"
          :rows="equityRows"
          @row-click="equityClick"
        >
        </q-table>
      </div>

      <div class="q-pa-md q-gutter-sm">
        <q-btn
          class="glossy"
          rounded
          label="Add Equity"
          color="primary"
          @click="persistentEquity = true"
        ></q-btn>

        <q-dialog
          v-model="persistentEquity"
          persistent
          transition-show="scale"
          transition-hide="scale"
        >
          <q-card class="bg-black text-white" style="width: 300px">
            <!-- Q-card-section is here for spacing on the title -->
            <q-card-section>
              <div class="text-h6 absolute-center"></div>
            </q-card-section>

            <q-card-section>
              <div class="text-h6 absolute-center">Equity Data</div>
            </q-card-section>

            <q-card-section>
              <form class="columns" action="" method="POST">
                <q-input
                  autogrow
                  filled
                  v-model="form.ticker"
                  label="Ticker"
                ></q-input>
                <q-input
                  autogrow
                  filled
                  v-model="form.shares"
                  label="Shares"
                  type="number"
                ></q-input>
                <q-input
                  autogrow
                  filled
                  v-model="form.equityPrice"
                  label="Price"
                  mask="$#.##"
                  input-class="text-left"
                  fill-mask="0"
                  reverse-fill-mask
                ></q-input>
              </form>
            </q-card-section>

            <q-card-section class="q-pt-none"> </q-card-section>

            <q-card-actions align="right" class="bg-black text-grey">
              <q-btn
                class="glossy"
                rounded
                color="primary"
                label="New Equity"
                @click="addEquity"
              ></q-btn>

              <q-popup-proxy>
                <q-banner>
                  <template v-slot:avatar>
                    <q-icon name="warning" color="primary"></q-icon>
                  </template>
                  Please enter a number of shares greater than 0 and a valid
                  symbol
                </q-banner>
              </q-popup-proxy>

              <q-btn flat label="Cancel" v-close-popup></q-btn>
            </q-card-actions>
          </q-card>
        </q-dialog>

        <!-- Edit Holdings Function -->

        <q-dialog
          v-model="persistentEquityEdit"
          persistent
          transition-show="scale"
          transition-hide="scale"
        >
          <q-card class="bg-black text-white" style="width: 300px">
            <q-card-section>
              <div class="text-h6 absolute center"></div>
            </q-card-section>

            <q-card-section>
              <div class="text-h6 absolute-center">Edit Holdings</div>
            </q-card-section>

            <q-card-section>
              <form class="columns" action="" method="POST">
                <q-input
                  disable
                  autogrow
                  filled
                  v-model="form.modticker"
                  label="Ticker"
                ></q-input>
                <q-input
                  autogrow
                  filled
                  v-model="form.modshares"
                  label="Shares"
                  type="number"
                ></q-input>
                <q-input
                  autogrow
                  filled
                  v-model="form.modprice"
                  label="Price"
                  mask="$#.##"
                  input-class="text-left"
                  fill-mask="0"
                  reverse-fill-mask
                ></q-input>
              </form>
            </q-card-section>

            <q-card-section class="q-pt-none"> </q-card-section>

            <q-card-actions align="right" class="bg-black text-grey">
              <q-btn
                class="glossy"
                rounded
                color="primary"
                label="Modify"
                @click="modifyEquity"
              ></q-btn>
              <q-btn
                class="glossy"
                rounded
                color="primary"
                label="Delete"
                @click="deleteEquity"
              ></q-btn>

              <q-popup-proxy>
                <q-banner>
                  <template v-slot:avatar>
                    <q-icon name="warning" color="primary"></q-icon>
                  </template>
                  Please enter a number of shares greater than 0 and a valid
                  symbol
                </q-banner>
              </q-popup-proxy>

              <q-btn flat label="Cancel" v-close-popup></q-btn>
            </q-card-actions>
          </q-card>
        </q-dialog>
      </div>

      <!-- Property Button with Dialog Box -->
      <q-btn
        class="glossy"
        rounded
        label="Add Property"
        color="primary"
        @click="persistentProperty = true"
      ></q-btn>

      <q-dialog
        v-model="persistentProperty"
        persistentProperty
        transition-show="scale"
        transition-hide="scale"
      >
        <q-card class="bg-black text-white" style="width: 300px">
          <!-- Q-card-section is here for spacing on the title -->
          <q-card-section>
            <div class="text-h6 absolute-center"></div>
          </q-card-section>

          <q-card-section>
            <div class="text-h6 absolute-center">New Property</div>
          </q-card-section>

          <q-card-section>
            <q-input
              autogrow
              filled
              v-model="form.description"
              label="Description"
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.price"
              label="Price"
              mask="$#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.lien"
              label="Lien"
              mask="$#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.rate"
              label="Rate"
              mask="##.##%"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.years"
              label="Years"
              type="number"
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.monthsLeft"
              label="Months Left"
              type="number"
            ></q-input>
            <q-input
              autogrow
              prefix="$"
              filled
              v-model="form.value"
              label="Value"
              type="number"
            ></q-input>
          </q-card-section>

          <q-card-section class="q-pt-none"> </q-card-section>

          <q-card-actions align="right" class="bg-black text-grey">
            <q-btn
              class="glossy"
              rounded
              color="primary"
              label="Add Property"
              @click="portfolioStore.importCurrentEquities"
            ></q-btn>

            <q-btn flat label="Cancel" v-close-popup></q-btn>
          </q-card-actions>
        </q-card>
      </q-dialog>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent, ref, onBeforeMount } from 'vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';
import { storeToRefs } from 'pinia';

// import EquityInput from './EquityInput.vue';

export default defineComponent({
  name: 'EquityComponent',
  props: {
    title: {
      type: String,
      required: true,
    },
  },
  setup() {
    let portfolioStore = usePortfolioStore();

    const { equities } = storeToRefs(portfolioStore);
    const { realEstate } = storeToRefs(portfolioStore);

    const { form } = storeToRefs(portfolioStore);

    const equitiesImported = ref(false);
    const persistentEquity = ref(false);
    const persistentProperty = ref(false);
    const persistentEquityEdit = ref(false);

    const columns = [
      {
        name: 'ticker',
        required: true,
        label: 'Ticker',
        align: 'left',
        field: (row) => row.ticker,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: 'shares',
        required: true,
        label: 'Shares',
        align: 'left',
        field: (row) => row.shares,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: 'price',
        required: true,
        label: 'Price',
        align: 'left',
        field: (row) => row.equityPrice,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: 'value',
        required: true,
        label: 'Value',
        align: 'left',
        field: (row) => row.value,
        format: (val) => `${val}`,
        sortable: true,
      },
    ];

    const { equityRows } = storeToRefs(portfolioStore);

    function equityClick(e: Event, row) {
      portfolioStore.form.modticker = row.ticker;
      portfolioStore.form.modshares = row.shares;
      portfolioStore.form.modprice = row.equityPrice;

      persistentEquityEdit.value = true;
    }

    onBeforeMount(() => {
      portfolioStore.importCurrentEquities();
    });

    return {
      equities,
      realEstate,
      portfolioStore,
      persistentEquity,
      persistentEquityEdit,
      persistentProperty,
      form,
      columns,
      equityRows,
      equitiesImported,
      equityClick,
      addEquity() {
        if (form.value.shares > 0 && typeof form.value.ticker == 'string') {
          portfolioStore.form.ticker = form.value.ticker;
          portfolioStore.form.shares = form.value.shares;
          portfolioStore.form.equityPrice = form.value.equityPrice;

          persistentEquity.value = false;
          portfolioStore.addEquityApi();

          portfolioStore.form.ticker = '';
          portfolioStore.form.shares = 0;
          portfolioStore.form.equityPrice = 0;
        } else {
        }
      },
      modifyEquity() {
        if (form.value.modshares > 0) {
          portfolioStore.form.modcurrentticker = form.value.modticker;
          portfolioStore.form.modcurrentshares = form.value.modshares;
          portfolioStore.form.modcurrentprice = form.value.modprice;

          portfolioStore.modifyEquity();

          persistentEquityEdit.value = false;
        }
      },
      deleteEquity() {
        portfolioStore.form.delcurrentticker = form.value.modticker;

        portfolioStore.deleteEquity();

        persistentEquityEdit.value = false;
      },
    };
  },
});
</script>
