<template>
  <q-page>
    <div style="text-align: center">
      <q-btn
        class="glossy"
        rounded
        color="primary"
        label="New Property"
        @click="newPropertyDialog"
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
              filled
              v-model="modelType"
              :options="options"
              label="Property Type"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) || 'Please select a property type',
              ]"
            />
            <q-input
              autogrow
              filled
              v-model="description"
              label="Street Address"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type property address here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="price"
              label="Purchase Price"
              prefix="$"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type what you paid for the property here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="lien"
              label="Mortgage"
              prefix="$"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type Your full mortgage amount here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="rate"
              label="Interest Rate"
              suffix="%"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type you interest rate here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="years"
              label="Years"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type the number of years on your mortgage here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="value"
              label="Market Value"
              prefix="$"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type the current market value of your property here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="ownership"
              label="Ownership Percentage"
              suffix="%"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type how much overship you have in the property here',
              ]"
            ></q-input>
            <q-input
              filled
              v-model="REdate"
              type="date"
              hint="Loan Closing Date"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) || 'Please type or select a date',
              ]"
            ></q-input>
          </form>
        </q-card-section>

        <q-card-section class="q-pt-none"> </q-card-section>

        <q-card-actions align="right" class="bg-black text-grey">
          <q-btn
            class="glossy"
            rounded
            color="primary"
            label="New Property"
            @click="addProperty(modelType as string)"
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

    <!-- Manage Properties Full Dialog Box -->
    <q-dialog
      v-model="dialog"
      persistent
      :maximized="maximizedToggle"
      transition-show="slide-up"
      transition-hide="slide-down"
    >
      <q-card class="bg-black text-primary">
        <q-banner dense inline-actions class="text-primary bg-black">
          <div style="font-size: 2rem; text-align: center">
            Property Dashboard
          </div>
          <template v-slot:action>
            <q-btn
              flat
              color="primary"
              label="Update Property"
              @click="persistentPropertyEdit = true"
            />
            <q-btn dense flat icon="close" v-close-popup>
              <q-tooltip class="bg-white text-primary">Close</q-tooltip>
            </q-btn>
          </template>
        </q-banner>

        <!-- Provides spacing between dashboard title and components -->
        <q-card-section>
          <div class="text-h6"></div>
        </q-card-section>

        <div class="flex row">
          <div class="col-2">
            <PropertySnapshot></PropertySnapshot>
          </div>

          <div class="col-9" style="margin-left: 4vw">
            <property-dashboard-component></property-dashboard-component>
          </div>

          <div class="flex col">
            <!-- Put apex chart in here for loan paydown -->
          </div>
        </div>
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
              filled
              v-model="modType"
              :options="options"
              label="Property Type"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) || 'Please select a property type',
              ]"
            />
            <q-input
              autogrow
              filled
              v-model="modDescription"
              label="Street Address"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type property address here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="modREprice"
              label="Purchase Price"
              prefix="$"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type what you originally paid for the property here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="modLien"
              label="Mortgage"
              prefix="$"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type Your full mortgage amount here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="modRate"
              label="Interest Rate"
              suffix="%"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type you interest rate here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="modYears"
              label="Years"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type the number of years on your mortgage here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="modValue"
              label="Market Value"
              prefix="$"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type the current market value of your property here',
              ]"
            ></q-input>
            <q-input
              autogrow
              filled
              type="number"
              v-model="modOwnership"
              label="Ownership Percentage"
              suffix="%"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) ||
                  'Please type how much overship you have in the property here',
              ]"
            ></q-input>
            <q-input
              filled
              v-model="modREDate"
              type="date"
              hint="Loan Closing Date"
              lazy-rules
              :rules="[
                (val) =>
                  (val && val.length > 0) || 'Please type or select a date',
              ]"
            ></q-input>
          </form>
        </q-card-section>

        <q-card-section class="q-pt-none"> </q-card-section>

        <q-card-actions align="right" class="bg-black text-grey">
          <q-btn
            class="glossy"
            rounded
            color="primary"
            label="Manage Property"
            @click="dialog = true"
          >
          </q-btn>
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
            @click="persistentSure = true"
          ></q-btn>

          <!-- Are you Sure Dialog -->
          <q-dialog
            v-model="persistentSure"
            persistent
            transition-show="scale"
            transition-hide="scale"
          >
            <q-card class="bg-black text-primary" style="width: 30vw">
              <q-card-section>
                <div class="text-h6" style="text-align: center">
                  Are you sure you want to delete this property?
                </div>
              </q-card-section>

              <q-card-actions align="right" class="bg-black text-primary">
                <q-btn
                  class="glossy"
                  rounded
                  color="primary"
                  label="Delete"
                  @click="deleteProperty"
                ></q-btn>
                <q-btn flat label="Cancel" v-close-popup />
              </q-card-actions>
            </q-card>
          </q-dialog>

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
import { Property } from 'src/stores/portfolio-store';
import PropertyDashboardComponent from './PropertyDashboardComponent.vue';
import PropertySnapshot from './PropertySnapshot.vue';

export default defineComponent({
  name: 'PropertyComponent',
  props: {
    title: {
      type: String,
      required: true,
    },
  },
  components: {
    PropertyDashboardComponent,
    PropertySnapshot,
  },
  setup() {
    let portfolioStore = usePortfolioStore();

    const { form } = storeToRefs(portfolioStore);

    const propertiesImported = ref(false);
    const persistentProperty = ref(false);
    const persistentPropertyEdit = ref(false);
    const persistentSure = ref(false);

    const realEstate = computed(() => portfolioStore.realEstate);
    const propertiesTotal = computed(() => portfolioStore.propertiesTotal);
    const selectedProperty = computed(
      () => portfolioStore.form.selectedProperty
    );

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
          ((row.equity as number) * (row.ownership as number)) / 100,
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

    function newPropertyDialog() {
      persistentProperty.value = true;
    }

    // Real Estate Forms
    const modelType = ref(null);
    const description = ref(null);
    const price = ref(null);
    const lien = ref(null);
    const rate = ref(null);
    const years = ref(null);
    const value = ref(null);
    const ownership = ref(null);
    const REdate = ref(null);
    const dialog = ref(false);
    // Modify Portion
    const modType = ref('');
    const modDescription = ref('');
    const modREprice = ref(0);
    const modLien = ref(0);
    const modRate = ref(0);
    const modYears = ref(0);
    const modValue = ref(0);
    const modOwnership = ref(0);
    const modREDate = ref('');

    return {
      // Property Form
      modelType,
      description,
      price,
      lien,
      rate,
      years,
      value,
      ownership,
      REdate,
      modType,
      modDescription,
      modREprice,
      modLien,
      modRate,
      modYears,
      modValue,
      modOwnership,
      modREDate,
      newPropertyDialog,

      portfolioStore,
      propertiesImported,
      persistentProperty,
      persistentPropertyEdit,
      realEstate,
      propertiesTotal,
      columns,
      propertyRows,
      propertyClick(e: Event, row: any) {
        portfolioStore.form.selectedProperty = new Property(
          ref(row.type),
          row.id,
          row.description,
          row.price,
          row.lien,
          row.rate,
          row.years,
          row.monthsLeft,
          row.value,
          row.ownership,
          row.date,
          row.tenants,
          row.buildingexpenses,
          row.buildingrevenues,
          row.ledger
        );

        dialog.value = true;

        // Conversion to "YYYY-MM-DD" required to use with Quasar date pickers
        const dd = row.date.slice(3, 5);
        const mm = row.date.slice(0, 2);
        const yyyy = row.date.slice(6, 10);
        const dateString = (yyyy + '-' + mm + '-' + dd) as string;

        modType.value = row.type;
        modDescription.value = row.description;
        modREprice.value = row.price;
        modLien.value = row.lien;
        modRate.value = row.rate;
        modYears.value = row.years;
        modValue.value = row.value;
        modOwnership.value = row.ownership;
        modREDate.value = dateString;

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
      },
      form,
      addProperty(modelType: string) {
        portfolioStore.form.type = modelType;
        portfolioStore.form.description = description.value;
        portfolioStore.form.price = price.value;
        portfolioStore.form.lien = lien.value;
        portfolioStore.form.rate = rate.value;
        portfolioStore.form.years = years.value;
        portfolioStore.form.value = value.value;
        portfolioStore.form.ownership = ownership.value;

        const dd = REdate.value.slice(8, 10);
        const mm = REdate.value.slice(5, 7);
        const yyyy = REdate.value.slice(0, 4);

        const dateString = mm + '-' + dd + '-' + yyyy;

        portfolioStore.form.REdate = dateString;

        persistentProperty.value = false;
        portfolioStore.addPropertyAPI();

        portfolioStore.form.description = '';
      },
      modifyProperty() {
        portfolioStore.form.modtype = modType.value;
        portfolioStore.form.moddescription = modDescription.value;
        portfolioStore.form.modREprice = modREprice.value;
        portfolioStore.form.modlien = modLien.value;
        portfolioStore.form.modrate = modRate.value;
        portfolioStore.form.modyears = modYears.value;
        portfolioStore.form.modvalue = modValue.value;
        portfolioStore.form.modOwnership = modOwnership.value;

        const dd = modREDate.value.slice(8, 10);
        const mm = modREDate.value.slice(5, 7);
        const yyyy = modREDate.value.slice(0, 4);

        const dateString = mm + '-' + dd + '-' + yyyy;

        portfolioStore.form.modREDate = dateString;

        portfolioStore.modifyProperty();

        persistentPropertyEdit.value = false;
      },
      options: ['Primary Residence', 'Second Home', 'Investment Property'],
      deleteProperty() {
        portfolioStore.form.delcurrentproperty = form.value.modid;

        portfolioStore.deleteProperty();

        persistentPropertyEdit.value = false;
      },
      dialog,
      maximizedToggle: ref(true),
      selectedProperty,
      persistentSure,
    };
  },
});
</script>
