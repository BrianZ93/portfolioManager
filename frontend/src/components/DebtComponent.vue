<template>
  <q-page>
    <div style="text-align: center">
      <q-btn
        class="glossy"
        rounded
        color="primary"
        label="New Debt"
        @click="persistentDebt = true"
      ></q-btn>
    </div>

    <q-dialog
      v-model="persistentDebt"
      persistent
      transition-show="scale"
      transition-hide="scale"
    >
      <q-card class="bg-black text-white" style="width: 400px">
        <!-- Q-card-section is here for spacing on the title -->
        <q-card-section>
          <div class="text-h6 absolute-center"></div>
        </q-card-section>

        <q-card-section>
          <div class="text-h6 absolute-center">Debt Data</div>
        </q-card-section>

        <q-card-section>
          <form class="columns" action="" method="POST">
            <q-select
              square
              filled
              v-model="modelType"
              :options="options"
              label="Debt Type"
            />
            <q-input
              autogrow
              filled
              v-model="form.debttitle"
              label="Title"
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.debtamount"
              label="Amount"
              prefix="$"
              mask="#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.debtrate"
              label="Interest Rate"
              mask="#.####"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.debtterm"
              label="Term (Years)"
              mask="#"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.debtdate"
              label="Date"
              @click="persistentDate = true"
            ></q-input>

            <div
              class="q-pa-md bg-grey-10 text-white"
              v-if="persistentDate == true"
            >
              <div class="q-gutter-md">
                <q-date v-model="date" color="primary" dark bordered />
                <q-btn
                  class="glossy"
                  rounded
                  color="primary"
                  label="Select Date"
                  @click="selectDate(date)"
                >
                </q-btn>
                <q-btn label="Cancel" @click="persistentDate = false"> </q-btn>
              </div>
            </div>
          </form>
        </q-card-section>

        <q-card-section class="q-pt-none"> </q-card-section>

        <q-card-actions align="right" class="bg-black text-grey">
          <q-btn
            class="glossy"
            rounded
            color="primary"
            label="New Debt"
            @click="addDebt(modelType)"
          ></q-btn>

          <!-- <q-popup-proxy>
            <q-banner>
              <template v-slot:avatar>
                <q-icon name="warning" color="primary"></q-icon>
              </template>
              Please enter a number of shares greater than 0 and a valid symbol
            </q-banner>
          </q-popup-proxy> -->

          <q-btn flat label="Cancel" v-close-popup></q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- Modify Dialog -->

    <q-dialog
      v-model="persistentDebtEdit"
      persistent
      transition-show="scale"
      transition-hide="scale"
    >
      <q-card class="bg-black text-white" style="width: 400px">
        <!-- Q-card-section is here for spacing on the title -->
        <q-card-section>
          <div class="text-h6 absolute-center"></div>
        </q-card-section>

        <q-card-section>
          <div class="text-h6 absolute-center">Debt Data</div>
        </q-card-section>

        <q-card-section>
          <form class="columns" action="" method="POST">
            <q-select
              square
              filled
              v-model="modelType"
              :options="options"
              label="Debt Type"
            />
            <q-input
              autogrow
              filled
              v-model="form.debtmodtitle"
              label="Title"
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.debtmodamount"
              label="Amount"
              prefix="$"
              mask="#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.debtmodrate"
              label="Interest Rate"
              mask="#.####"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.debtmodterm"
              label="Term (Years)"
              mask="#"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.debtmoddate"
              label="Date"
              @click="persistentDate = true"
            ></q-input>

            <div
              class="q-pa-md bg-grey-10 text-white"
              v-if="persistentDate == true"
            >
              <div class="q-gutter-md">
                <q-date v-model="date" color="primary" dark bordered />
                <q-btn
                  class="glossy"
                  rounded
                  color="primary"
                  label="Select Date"
                  @click="selectDate(date)"
                >
                </q-btn>
                <q-btn label="Cancel" @click="persistentDate = false"> </q-btn>
              </div>
            </div>
          </form>
        </q-card-section>

        <q-card-section class="q-pt-none"> </q-card-section>

        <q-card-actions align="right" class="bg-black text-grey">
          <q-btn
            class="glossy"
            rounded
            color="primary"
            label="Modify"
            @click="modifyDebt"
          ></q-btn>
          <q-btn
            class="glossy"
            rounded
            color="primary"
            label="Delete"
            @click="deleteDebt"
          ></q-btn>

          <!-- <q-popup-proxy>
            <q-banner>
              <template v-slot:avatar>
                <q-icon name="warning" color="primary"></q-icon>
              </template>
              Please enter a number of shares greater than 0 and a valid symbol
            </q-banner>
          </q-popup-proxy> -->

          <q-btn flat label="Cancel" v-close-popup></q-btn>
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- Debt Table -->
    <div class="row">
      <div class="col q-pa-md">
        <q-table
          :title="
            'Debt - ' +
            debtsTotal
              .toLocaleString('en-US', { style: 'currency', currency: 'USD' })
              .toString()
          "
          :columns="columns"
          :rows="Debts"
          :key="Debts"
          @row-click="DebtClick"
        >
        </q-table>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
// TODO - Change the input types and change the way percentages are reflected

import { defineComponent, ref, computed, onBeforeMount } from 'vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';
import { storeToRefs } from 'pinia';

export default defineComponent({
  name: 'DebtComponent',
  setup() {
    let portfolioStore = usePortfolioStore();

    const persistentDebt = ref(false);
    const persistentDebtEdit = ref(false);
    const persistentDate = ref(false);

    const { Debts } = storeToRefs(portfolioStore);
    const { form } = storeToRefs(portfolioStore);

    onBeforeMount(() => {
      portfolioStore.importCurrentEquities();
      portfolioStore.importCurrentProperties();
      portfolioStore.importCurrentDebts();
    });

    const debtsTotal = computed(() => portfolioStore.debtsTotal);

    // Property Columns
    const columns = [
      {
        name: 'title',
        required: true,
        label: 'Title',
        align: 'left',
        field: (row: Record<string, unknown>) => row.title,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'type',
        required: true,
        label: 'Type',
        align: 'left',
        field: (row: Record<string, unknown>) => row.type,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'original amount',
        required: true,
        label: 'Original Amount',
        align: 'left',
        field: (row: Record<string, unknown>) => row.amount,
        format: (val: number) =>
          val.toLocaleString('en-US', { style: 'currency', currency: 'USD' }),
        sortable: true,
      },
      {
        name: 'current',
        required: true,
        label: 'Current Amount',
        align: 'left',
        field: (row: Record<string, unknown>) => row.currentBalance,
        format: (val: number) =>
          val.toLocaleString('en-US', { style: 'currency', currency: 'USD' }),
        sortable: true,
      },
      {
        name: 'rate',
        required: true,
        label: 'Rate',
        align: 'left',
        field: (row: Record<string, unknown>) => row.rate,
        format: (val: number) => `${val}` + '%',
        sortable: true,
      },
      {
        name: 'term',
        required: true,
        label: 'Term (Years)',
        align: 'left',
        field: (row: Record<string, unknown>) => row.term,
        format: (val: number) => `${val}`,
        sortable: true,
      },
      {
        name: 'months left',
        required: true,
        label: 'Months Remaining',
        align: 'left',
        field: (row: Record<string, unknown>) => row.monthsLeft,
        format: (val: number) => `${val}`,
        sortable: true,
      },
      {
        name: 'start date',
        required: true,
        label: 'Start Date',
        align: 'left',
        field: (row: Record<string, unknown>) => row.date,
        format: (val: number) => `${val}`,
        sortable: true,
      },
    ];

    const modelType = ref('');

    return {
      persistentDebt,
      persistentDate,
      persistentDebtEdit,
      Debts,
      debtsTotal,
      columns,
      modelType,
      options: ['Installment', 'Revolving', 'Lease'],
      form,
      date: ref('2022/10/01'),
      selectDate(date: string) {
        // Date Constructor (Original Quasar Object has YYYY/MM/DD Format)
        const year = date.slice(0, 4);
        const month = date.slice(5, 7);
        const day = date.slice(8, 10);

        const datestring = month + '/' + day + '/' + year;

        portfolioStore.form.debtdate = datestring;
        portfolioStore.form.debtmoddate = datestring;

        persistentDate.value = false;
      },
      addDebt(modelType: string) {
        portfolioStore.form.debttitle = form.value.debttitle;
        portfolioStore.form.debttype = modelType;
        portfolioStore.form.debtamount = form.value.debtamount;
        portfolioStore.form.debtrate = form.value.debtrate;
        portfolioStore.form.debtterm = form.value.debtterm;
        portfolioStore.form.debtdate = form.value.debtdate;

        persistentDebt.value = false;
        portfolioStore.addDebtAPI();

        portfolioStore.form.debttitle = '';
      },
      DebtClick(e: Event, row: any) {

        portfolioStore.form.debtmodid = row.id;
        portfolioStore.form.debtmodtitle = row.title;
        portfolioStore.form.debtmodtype = ref(row.type);
        portfolioStore.form.debtmodamount = row.amount;
        portfolioStore.form.debtmodrate = row.rate;
        portfolioStore.form.debtmodterm = row.term;
        portfolioStore.form.debtmoddate = row.date;
        modelType.value = row.type;

        persistentDebtEdit.value = true;
        persistentDate.value = false;
      },
      modifyDebt() {
        portfolioStore.form.debtmodid = form.value.debtmodid;
        portfolioStore.form.debtmodtitle = form.value.debtmodtitle;
        portfolioStore.form.debtmodtype = modelType.value;
        portfolioStore.form.debtmodamount = form.value.debtmodamount;
        portfolioStore.form.debtmodrate = form.value.debtmodrate;
        portfolioStore.form.debtmodterm = form.value.debtmodterm;
        portfolioStore.form.debtmoddate = form.value.debtmoddate;

        portfolioStore.modifyDebt();

        persistentDebtEdit.value = false;
      },
      deleteDebt() {
        portfolioStore.form.delcurrentdebt = form.value.debtmodid;

        portfolioStore.deleteDebt();

        persistentDebtEdit.value = false;
      },
    };
  },
});
</script>
