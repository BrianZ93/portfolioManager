<template>
  <q-card class="my-card">
    <q-item>
      <q-item-section class="text-h6 text-primary" style="text-align: center">
        Property Snapshot
      </q-item-section>
    </q-item>
    <q-separator></q-separator>
    <q-list>
      <q-item clickable>
        <q-item-section avatar>
          <q-icon color="primary" name="home" />
        </q-item-section>

        <q-item-section>
          <q-item-label> {{ selectedProperty.description }}</q-item-label>
          <q-item-label caption>{{ selectedProperty.type }}</q-item-label>
        </q-item-section>
      </q-item>
      <q-item clickable>
        <q-item-section avatar>
          <q-icon color="green" name="paid" />
        </q-item-section>

        <q-item-section>
          <q-item-label>Current Value</q-item-label>
          <q-item-label caption>{{
            selectedProperty.value.toLocaleString('en-US', {
              style: 'currency',
              currency: 'USD',
            })
          }}</q-item-label>
        </q-item-section>
      </q-item>

      <q-item clickable>
        <q-item-section avatar>
          <q-icon color="red" name="payments" />
        </q-item-section>

        <q-item-section>
          <q-item-label>Outstanding Liens</q-item-label>
          <q-item-label caption>{{
            selectedProperty.lien.toLocaleString('en-US', {
              style: 'currency',
              currency: 'USD',
            })
          }}</q-item-label>
        </q-item-section>
      </q-item>

      <q-item clickable>
        <q-item-section avatar>
          <q-icon
            v-if="selectedProperty.value - selectedProperty.lien < 0"
            color="red"
            name="wallet"
          />
          <q-icon
            v-if="selectedProperty.value - selectedProperty.lien > -0.01"
            color="green"
            name="wallet"
          />
        </q-item-section>

        <q-item-section>
          <q-item-label>Equity</q-item-label>
          <q-item-label caption>{{
            (selectedProperty.value - selectedProperty.lien).toLocaleString(
              'en-us',
              {
                style: 'currency',
                currency: 'USD',
              }
            )
          }}</q-item-label>
        </q-item-section>
      </q-item>
    </q-list>

    <q-item clickable @click="persistentTenants = true">
      <q-item-section avatar>
        <q-icon color="amber" name="perm_identity" />
      </q-item-section>

      <q-item-section>
        <q-item-label>Add Tenant</q-item-label>
      </q-item-section>
    </q-item>

    <q-item>
      <q-item-section class="text-h6 text-primary" style="text-align: center">
        Building Expenses & Other Revenues
      </q-item-section>
    </q-item>
    <q-separator></q-separator>
    <q-item clickable @click="persistentRevenues = true">
      <q-item-section avatar>
        <q-icon color="green" name="attach_money" />
      </q-item-section>

      <q-item-section>
        <q-item-label>Revenues</q-item-label>
        <q-item-label caption>Add Caption Later</q-item-label>
      </q-item-section>
    </q-item>
    <q-item clickable @click="persistentExpenses = true">
      <q-item-section avatar>
        <q-icon color="red" name="payments" />
      </q-item-section>

      <q-item-section>
        <q-item-label>Expenses</q-item-label>
        <q-item-label caption>Add Caption Later</q-item-label>
      </q-item-section>
    </q-item>
  </q-card>

  <q-dialog
    v-model="persistentTenants"
    persistent
    transition-show="scale"
    transition-hide="scale"
  >
    <q-card class="text-primary" style="width: 30vw">
      <q-card-section>
        <div class="text-h6 absolute-center">Tenants</div>
      </q-card-section>

      <q-card-section>
        <q-list>
          <q-item
            v-for="tenant in form.selectedProperty.tenants"
            :key="tenant && selectedProperty"
            :value="tenant"
          >
            <q-item-label>{{ tenant.name }}</q-item-label>
          </q-item>

          <q-expansion-item icon="perm_identity" label="Add Tenant">
            <q-form
              @submit="onSubmit(selectedProperty)"
              @reset="onTenantReset"
              class="q-gutter-md"
            >
              <q-input
                filled
                v-model="name"
                label="Tenant Name"
                hint="Name and surname"
                lazy-rules
                :rules="[
                  (val) => (val && val.length > 0) || 'Please type something',
                ]"
              />

              <q-input
                filled
                v-model="unit"
                label="Tenant Unit Number"
                hint="If no units type anything here"
                lazy-rules
                :rules="[
                  (val) => (val && val.length > 0) || 'Please type something',
                ]"
              />

              <q-input
                filled
                v-model="start"
                type="date"
                hint="Lease Start Date"
                lazy-rules
                :rules="[
                  (val) => (val && val.length > 0) || 'Please type something',
                ]"
              />

              <q-input
                filled
                v-model="end"
                type="date"
                hint="Lease End Date"
                lazy-rules
                :rules="[
                  (val) => (val && val.length > 0) || 'Please type something',
                ]"
              />

              <q-toggle
                v-model="currentTenant"
                label="This tenant currently lives here"
              />

              <div>
                <q-btn label="Submit" type="submit" color="primary" />
                <q-btn
                  label="Reset"
                  type="reset"
                  color="primary"
                  flat
                  class="q-ml-sm"
                />
              </div>
            </q-form>
          </q-expansion-item>
        </q-list>
      </q-card-section>

      <q-card-actions align="right" class="text-primary">
        <q-btn flat label="OK" v-close-popup />
      </q-card-actions>
    </q-card>
  </q-dialog>

  <q-dialog
    v-model="persistentExpenses"
    persistent
    transition-show="scale"
    transition-hide="scale"
  >
    <q-card class="text-primary" style="width: 300px">
      <q-card-section>
        <div class="text-h6" style="text-align: center">
          Add Building Expense
        </div>
      </q-card-section>

      <q-card-section class="q-pt-none" style="text-align: center">
        Enter an expense for {{ selectedProperty.description }}
      </q-card-section>

      <q-form
        @submit="onExpenseSubmit(selectedProperty)"
        @reset="onReset"
        class="q-gutter-md"
      >
        <q-input
          autogrow
          filled
          v-model="expense"
          label="Description"
          hint="Brief description of your expense"
          lazy-rules
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
        />

        <q-input
          autogrow
          filled
          type="number"
          v-model="expenseamount"
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
          v-model="expensedate"
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

  <q-dialog
    v-model="persistentRevenues"
    persistent
    transition-show="scale"
    transition-hide="scale"
  >
    <q-card class="text-primary" style="width: 300px">
      <q-card-section>
        <div class="text-h6" style="text-align: center">
          Add Building Revenue
        </div>
      </q-card-section>

      <q-card-section class="q-pt-none" style="text-align: center">
        Enter a revenue item for {{ selectedProperty.description }}
      </q-card-section>

      <q-form
        @submit="onRevenueSubmit(selectedProperty)"
        @reset="onRevenueReset"
        class="q-gutter-md"
      >
        <q-input
          autogrow
          filled
          v-model="revenue"
          label="Description"
          hint="Brief description of your revenue"
          lazy-rules
          :rules="[(val) => (val && val.length > 0) || 'Please type something']"
        />

        <q-input
          autogrow
          filled
          type="number"
          v-model="revamount"
          label="Amount"
          prefix="$"
          lazy-rules
          :rules="[
            (val) =>
              (val !== null && val !== '') || 'Please type your revenue amount',
          ]"
        />

        <q-input
          filled
          v-model="revdate"
          type="date"
          hint="Date of revenue"
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
import { usePortfolioStore } from 'src/stores/portfolio-store';
import { defineComponent, computed, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { Property, Tenant, Expense, Revenue } from 'src/stores/portfolio-store';

export default defineComponent({
  name: 'PropertySnapshot',
  setup() {
    let portfolioStore = usePortfolioStore();

    const { form } = storeToRefs(portfolioStore);

    const selectedProperty = computed(
      () => portfolioStore.form.selectedProperty
    );

    const persistentTenants = ref(false);
    const persistentExpenses = ref(false);
    const persistentRevenues = ref(false);

    // Tenant Form
    const name = ref(null);
    const unit = ref(null);
    const start = ref(null);
    const end = ref(null);
    const currentTenant = ref(false);

    // Expense Form
    const expense = ref(null);
    const expenseamount = ref(null);
    const expensedate = ref(null);

    // Revenue Form
    const revenue = ref(null);
    const revamount = ref(null);
    const revdate = ref(null);

    function onSubmit(selectedProperty: Property) {
      portfolioStore.form.newTenant = new Tenant(
        name.value,
        start.value,
        end.value,
        [] as Array<Expense>,
        [] as Array<Revenue>,
        unit.value,
        currentTenant.value,
        selectedProperty.id,
        0
      );

      portfolioStore.addTenant();
    }

    function onExpenseSubmit(selectedProperty: Property) {
      const amount = Number(expenseamount.value);
      portfolioStore.form.newExpense = new Expense(
        selectedProperty.id,
        0,
        expense.value,
        amount,
        expensedate.value,
        0
      );

      persistentExpenses.value = false;

      portfolioStore.addExpense();
    }

    function onRevenueSubmit(selectedProperty: Property) {
      const amount = Number(revamount.value);
      portfolioStore.form.newRevenue = new Revenue(
        selectedProperty.id,
        0,
        revenue.value,
        amount,
        revdate.value,
        0
      );

      persistentRevenues.value = false;

      portfolioStore.addRevenue();
    }

    return {
      selectedProperty,
      persistentTenants,
      persistentExpenses,
      persistentRevenues,
      form,
      onSubmit,
      onExpenseSubmit,
      onRevenueSubmit,
      name,
      unit,
      start,
      end,
      currentTenant,
      onReset() {
        expense.value = null;
        expenseamount.value = null;
        expensedate.value = null;
      },
      onTenantReset() {
        name.value = null;
        unit.value = null;
        start.value = null;
        end.value = null;
      },
      onRevenueReset() {
        revenue.value = null;
        revamount.value = null;
        revdate.value = null;
      },
      expense,
      expenseamount,
      expensedate,
      revenue,
      revamount,
      revdate,
    };
  },
});
</script>

<style lang="sass" scoped>
.my-card
  width: 100%
  max-width: 15vw
  margin-left: 1vw
</style>
