<template>
  <q-page>
    <div style="text-align: center">
      <q-btn
        class="glossy"
        rounded
        color="primary"
        label="New Property"
        @click="persistentProperty = true"
      ></q-btn>
    </div>

    <q-dialog
      v-model="persistentProperty"
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
          <div class="text-h6 absolute-center">Property Data</div>
        </q-card-section>

        <q-card-section>
          <form class="columns" action="" method="POST">
            <q-select
              square
              filled
              v-model="modelType"
              :options="options"
              label="Property Type"
            />
            <q-input
              autogrow
              filled
              v-model="form.description"
              label="Street Address"
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.price"
              label="Purchase Price"
              prefix="$"
              mask="#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.lien"
              label="Loan Amount"
              mask="$#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.rate"
              label="Interest Rate"
              mask="#.####"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.years"
              label="Term (years)"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.monthsLeft"
              label="Term (months)"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
              disable
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.value"
              label="Current Market Value"
              mask="$#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.ownership"
              label="Ownership %"
              prefix="$"
              mask="#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.REdate"
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
            label="New Property"
            @click="addProperty(modelType)"
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
      v-model="persistentPropertyEdit"
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
          <div class="text-h6 absolute-center">Property Data</div>
        </q-card-section>

        <q-card-section>
          <form class="columns" action="" method="POST">
            <q-select
              square
              filled
              v-model="form.modtype"
              :options="options"
              label="Property Type"
            />
            <q-input
              autogrow
              filled
              v-model="form.moddescription"
              label="Street Address"
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.modREprice"
              label="Purchase Price"
              mask="$#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.modlien"
              label="Loan Amount"
              mask="$#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.modrate"
              label="Interest Rate"
              mask="#.####"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.modyears"
              label="Term (years)"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.modmonthsLeft"
              label="Term (months)"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
              disable
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.modvalue"
              label="Current Market Value"
              mask="$#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.modOwnership"
              label="Ownership %"
              mask="#.##"
              input-class="text-left"
              fill-mask="0"
              reverse-fill-mask
            ></q-input>
            <q-input
              autogrow
              filled
              v-model="form.modREDate"
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
            @click="modifyProperty"
          ></q-btn>
          <q-btn
            class="glossy"
            rounded
            color="primary"
            label="Delete"
            @click="deleteProperty"
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

    <div class="row">
      <div class="col q-pa-md">
        <q-table
          :title="
            'Properties - ' +
            propertiesTotal.toLocaleString('en-US', {
              style: 'currency',
              currency: 'USD',
            })
          "
          :columns="columns"
          :rows="propertyRows"
          :key="propertyRows"
          @row-click="propertyClick"
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
  name: 'PropertyComponent',
  props: {
    title: {
      type: String,
      required: true,
    },
  },
  setup() {
    let portfolioStore = usePortfolioStore();

    const { form } = storeToRefs(portfolioStore);

    const propertiesImported = ref(false);
    const persistentProperty = ref(false);
    const persistentPropertyEdit = ref(false);
    const persistentDate = ref(false);

    const realEstate = computed(() => portfolioStore.realEstate);
    const propertiesTotal = computed(() => portfolioStore.propertiesTotal);

    // Property Columns
    const columns = [
      {
        name: 'address',
        required: true,
        label: 'Street Address',
        align: 'left',
        field: (row: Record<string, unknown>) => row.description,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'type',
        required: true,
        label: 'Type',
        align: 'left',
        field: (row: Record<string, unknown>) => row.type,
        format: (val: number) => `${val}`,
        sortable: true,
      },
      {
        name: 'value',
        required: true,
        label: 'Market Value',
        align: 'left',
        field: (row: Record<string, unknown>) => row.value,
        format: (val: number) =>
          val.toLocaleString('en-US', { style: 'currency', currency: 'USD' }),
        sortable: true,
      },
      {
        name: 'original loan balance',
        required: true,
        label: 'Original Loan Balance',
        align: 'left',
        field: (row: Record<string, unknown>) => row.lien,
        format: (val: number) =>
          val.toLocaleString('en-US', { style: 'currency', currency: 'USD' }),
        sortable: true,
      },
      {
        name: 'current loan balance',
        required: true,
        label: 'Current Loan Balance',
        align: 'left',
        field: (row: Record<string, unknown>) => row.currentBalance,
        format: (val: number) =>
          val.toLocaleString('en-US', { style: 'currency', currency: 'USD' }),
        sortable: true,
      },
      {
        name: 'equity',
        required: true,
        label: 'Equity',
        align: 'left',
        field: (row: Record<string, unknown>) => row.equity,
        format: (val: number) =>
          val.toLocaleString('en-US', { style: 'currency', currency: 'USD' }),
        sortable: true,
      },
      {
        name: 'ownership',
        required: true,
        label: 'Ownership',
        align: 'left',
        field: (row: Record<string, unknown>) =>
          (row.equity * row.ownership) / 100,
        format: (val: number) =>
          val.toLocaleString('en-US', { style: 'currency', currency: 'USD' }),
        sortable: true,
      },
      {
        name: 'loan start date',
        required: true,
        label: 'Loan Start Date',
        align: 'left',
        field: (row: Record<string, unknown>) => row.date,
        format: (val: string) => `${val}`,
        sortable: true,
      },
    ];

    const { propertyRows } = storeToRefs(portfolioStore);

    onBeforeMount(() => {
      portfolioStore.importCurrentEquities();
      portfolioStore.importCurrentProperties();
      portfolioStore.importCurrentDebts();
    });

    return {
      portfolioStore,
      propertiesImported,
      persistentProperty,
      persistentPropertyEdit,
      realEstate,
      propertiesTotal,
      columns,
      propertyRows,
      propertyClick(e: Event, row: any) {
        console.log(row);

        portfolioStore.form.modid = row.id;
        portfolioStore.form.modtype = ref(row.type);
        portfolioStore.form.moddescription = row.description;
        portfolioStore.form.modREprice = row.price;
        portfolioStore.form.modlien = row.lien;
        portfolioStore.form.modrate = row.rate;
        portfolioStore.form.modyears = row.years;
        portfolioStore.form.modmonthsLeft = row.monthsLeft;
        portfolioStore.form.modvalue = row.value;
        portfolioStore.form.modOwnership = row.ownership;
        portfolioStore.form.modREDate = row.date;

        persistentPropertyEdit.value = true;
        persistentDate.value = false;
      },
      form,
      addProperty(modelType: string) {
        portfolioStore.form.type = modelType;
        portfolioStore.form.description = form.value.description;
        portfolioStore.form.price = form.value.price;
        portfolioStore.form.lien = form.value.lien;
        portfolioStore.form.rate = form.value.rate;
        portfolioStore.form.years = form.value.years;
        portfolioStore.form.value = form.value.value;
        portfolioStore.form.ownership = form.value.ownership;
        portfolioStore.form.REdate = form.value.REdate;

        persistentProperty.value = false;
        portfolioStore.addPropertyAPI();

        portfolioStore.form.description = '';
      },
      modifyProperty() {
        portfolioStore.form.modtype = form.value.modtype;
        portfolioStore.form.moddescription = form.value.moddescription;
        portfolioStore.form.modREprice = form.value.modREprice;
        portfolioStore.form.modlien = form.value.modlien;
        portfolioStore.form.modrate = form.value.modrate;
        portfolioStore.form.modyears = form.value.modyears;
        portfolioStore.form.modmonthsLeft = form.value.modmonthsLeft;
        portfolioStore.form.modvalue = form.value.modvalue;
        portfolioStore.form.modOwnership = form.value.modOwnership;
        portfolioStore.form.modREDate = form.value.modREDate;

        portfolioStore.modifyProperty();

        persistentPropertyEdit.value = false;
      },
      modelType: ref(''),
      options: ['Primary Residence', 'Second Home', 'Investment Property'],
      deleteProperty() {
        portfolioStore.form.delcurrentproperty = form.value.modid;

        portfolioStore.deleteProperty();

        persistentPropertyEdit.value = false;
      },
      persistentDate,
      date: ref('2022/10/01'),
      selectDate(date: string) {
        // Date Constructor (Original Quasar Object has YYYY/MM/DD Format)
        const year = date.slice(0, 4);
        const month = date.slice(5, 7);
        const day = date.slice(8, 10);

        const datestring = month + '/' + day + '/' + year;

        portfolioStore.form.REdate = datestring;
        portfolioStore.form.modREDate = datestring;

        persistentDate.value = false;
      },
    };
  },
});
</script>
