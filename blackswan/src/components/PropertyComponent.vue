<template>
  <q-page>
    <h1 style="text-align: center">{{ title }}</h1>

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
      <q-card class="bg-black text-white" style="width: 300px">
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
              mask="$#.##"
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
              mask="#.##%"
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
          </form>
        </q-card-section>

        <q-card-section class="q-pt-none"> </q-card-section>

        <q-card-actions align="right" class="bg-black text-grey">
          <q-btn
            class="glossy"
            rounded
            color="primary"
            label="New Property"
            @click="addProperty"
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
      <div class="col-xs-12 col-sm-6 col-md-4">col</div>
      <div class="col-xs-12 col-sm-6 col-md-4">col</div>
      <div class="col-xs-12 col-sm-6 col-md-4">col</div>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';
import { storeToRefs } from 'pinia';
// import { PropertySummary } from './models';

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

    return {
      portfolioStore,
      propertiesImported,
      persistentProperty,
      persistentPropertyEdit,
      form,
      addProperty() {
        console.log('attempted');
        console.log('it worked');
        portfolioStore.form.type = form.value.type;
        portfolioStore.form.description = form.value.description;
        portfolioStore.form.price = form.value.price;
        portfolioStore.form.lien = form.value.lien;
        portfolioStore.form.rate = form.value.rate;
        portfolioStore.form.years = form.value.years;
        portfolioStore.form.value = form.value.value;

        persistentProperty.value = false;
        portfolioStore.addPropertyAPI();

        portfolioStore.form.description = '';
        // portfolioStore.form.price = 0;
        // portfolioStore.form.lien = 0;
        // portfolioStore.form.rate = 0;
        // portfolioStore.form.years = 0;
        // portfolioStore.form.value = 0;
      },
      modelType: ref(null),
      options: ['Primary Residence', 'Second Home', 'Investment Property'],
    };
  },
});
</script>
