<template>
  <q-card>
    <q-table
      dense
      title="General Ledger"
      :columns="columns"
      :rows="form.selectedProperty.ledger"
      :key="form.selectedProperty.ledger"
      @row-click="itemClick"
    ></q-table>
  </q-card>

  <!-- Ledger Modify Dialog -->

  <q-dialog
    v-model="persistentLedgerEdit"
    persistent
    transition-show="scale"
    transition-hide="scale"
  >
    <q-card class="text-primary" style="width: 300px">
      <q-card-section>
        <div class="text-h6" style="text-align: center">Edit Ledger Item</div>
      </q-card-section>

      <q-card-section class="q-pt-none" style="text-align: center">
        Edit Entry Below
      </q-card-section>

      <q-form
        @submit="onLedgerModifySubmit(form.selectedProperty)"
        @reset="onLedgerModifyReset"
        class="q-gutter-md"
      >
        <q-input autogrow filled v-model="modtype" label="Type" disable />
        <q-input
          autogrow
          filled
          v-model="moddescription"
          label="Description"
          hint="Brief description of your item"
          lazy-rules
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
        />

        <q-input
          autogrow
          filled
          type="number"
          v-model="modamount"
          label="Amount"
          prefix="$"
          lazy-rules
          :rules="[
            (val) =>
              (val !== null && val !== '') || 'Please type your expense amount',
          ]"
        />

        <q-input
          filled
          v-model="moddate"
          type="date"
          hint="Date of expense"
          lazy-rules
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
        />

        <div style="margin-bottom: 1vh">
          <q-btn label="Submit" type="submit" color="primary" />
          <q-btn
            label="Reset"
            type="reset"
            color="primary"
            flat
            class="q-ml-sm"
          />
          <q-btn flat label="Cancel" v-close-popup />
        </div>
      </q-form>
    </q-card>
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';
import { storeToRefs } from 'pinia';
import { Expense, Revenue } from 'src/stores/portfolio-store';

export default defineComponent({
  name: 'generalLedger',
  setup() {
    let portfolioStore = usePortfolioStore();

    const { form } = storeToRefs(portfolioStore);

    const persistentLedgerEdit = ref(false);

    const modtype = ref(null);
    const moddescription = ref(null);
    const modamount = ref(null);
    const moddate = ref(null);

    function itemClick(e: Event, row: any) {
      console.log(row);
      modtype.value = row.type;
      moddescription.value = row.description;
      modamount.value = row.amount;
      moddate.value = row.date;

      portfolioStore.form.modLedgerId = row.id;
      portfolioStore.form.modLedgerTenantId = row.tenantid;
      portfolioStore.form.modLedgerSubId = row.subid;

      persistentLedgerEdit.value = true;
    }

    function onLedgerModifyReset() {
      moddescription.value = null;
      modamount.value = null;
      moddate.value = null;
    }

    function onLedgerModifySubmit() {
      if (modtype.value == 'Expense') {
        const newExpense = new Expense(
          portfolioStore.form.modLedgerId,
          portfolioStore.form.modLedgerTenantId,
          moddescription.value,
          modamount.value,
          moddate.value,
          portfolioStore.form.modLedgerSubId
        );

        console.log(newExpense);

        portfolioStore.modifyExpense(newExpense);
      } else {
        const newRevenue = new Revenue(
          portfolioStore.form.modLedgerId,
          portfolioStore.form.modLedgerTenantId,
          moddescription.value,
          modamount.value,
          moddate.value,
          portfolioStore.form.modLedgerSubId
        );

        portfolioStore.modifyRevenue(newRevenue);
      }
    }

    // Ledger Columns
    const columns = [
      {
        name: 'date',
        required: true,
        label: 'Date',
        align: 'left',
        field: (row: Record<string, unknown>) => row.date,
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
        name: 'description',
        required: true,
        label: 'Description',
        align: 'left',
        field: (row: Record<string, unknown>) => row.description,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'amount',
        required: true,
        label: 'Amount',
        align: 'left',
        field: (row: Record<string, unknown>) => row.amount,
        format: (val: number) =>
          val.toLocaleString('en-US', { style: 'currency', currency: 'USD' }),
        sortable: true,
      },
    ];

    return {
      form,
      columns,
      persistentLedgerEdit,
      itemClick,
      modtype,
      moddescription,
      modamount,
      moddate,
      onLedgerModifyReset,
      onLedgerModifySubmit,
    };
  },
});
</script>
