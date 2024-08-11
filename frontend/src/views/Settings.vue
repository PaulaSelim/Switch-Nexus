<script setup>
import { ref, onMounted } from "vue";
import { FilterMatchMode } from "@primevue/core/api";
import { useToast } from "primevue/usetoast";
import MySwitchService from "../service/MySwitchService.js";
import Dialog from "primevue/dialog";
import Button from "primevue/button";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Toolbar from "primevue/toolbar";
import InputText from "primevue/inputtext";
import IconField from "primevue/iconfield";
import InputIcon from "primevue/inputicon";
import { useRouter } from "vue-router";
import MegaMenu from "primevue/megamenu";

const router = useRouter();

function navigateToPage(home) {
  router.push("/");
}

onMounted(() => {
  MySwitchService.getMySwitches().then((data) => (myswitches.value = data));
});

const toast = useToast();
const dt = ref();
const myswitches = ref();
const myswitchDialog = ref(false);
const deleteMySwitchDialog = ref(false);
const deleteMySwitchesDialog = ref(false);
const myswitch = ref({});
const selectedMySwitches = ref();
const filters = ref({
  global: { value: null, matchMode: FilterMatchMode.CONTAINS },
});
const submitted = ref(false);

const openNew = () => {
  myswitch.value = {};
  submitted.value = false;
  myswitchDialog.value = true;
};
const hideDialog = () => {
  myswitchDialog.value = false;
  submitted.value = false;
};
const saveMySwitch = () => {
  submitted.value = true;

  if (myswitch?.value.name?.trim()) {
    if (myswitch.value.id) {
      myswitch.value.inventoryStatus = myswitch.value.inventoryStatus.value
        ? myswitch.value.inventoryStatus.value
        : myswitch.value.inventoryStatus;
      myswitches.value[findIndexById(myswitch.value.id)] = myswitch.value;
      toast.add({
        severity: "success",
        summary: "Successful",
        detail: "MySwitch Updated",
        life: 3000,
      });
    } else {
      myswitch.value.id = createId();
      myswitch.value.code = createId();
      myswitch.value.inventoryStatus = myswitch.value.inventoryStatus
        ? myswitch.value.inventoryStatus.value
        : "INSTOCK";
      myswitches.value.push(myswitch.value);
      toast.add({
        severity: "success",
        summary: "Successful",
        detail: "MySwitch Created",
        life: 3000,
      });
    }

    myswitchDialog.value = false;
    myswitch.value = {};
  }
};
const editMySwitch = (prod) => {
  myswitch.value = { ...prod };
  myswitchDialog.value = true;
};
const confirmDeleteMySwitch = (prod) => {
  myswitch.value = prod;
  deleteMySwitchDialog.value = true;
};
const deleteMySwitch = () => {
  myswitches.value = myswitches.value.filter((val) => val.id !== myswitch.value.id);
  deleteMySwitchDialog.value = false;
  myswitch.value = {};
  toast.add({
    severity: "success",
    summary: "Successful",
    detail: "MySwitch Deleted",
    life: 3000,
  });
};
const findIndexById = (id) => {
  let index = -1;
  for (let i = 0; i < myswitches.value.length; i++) {
    if (myswitches.value[i].id === id) {
      index = i;
      break;
    }
  }

  return index;
};

const createId = () => {
  let id = "";
  var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  for (var i = 0; i < 5; i++) {
    id += chars.charAt(Math.floor(Math.random() * chars.length));
  }
  return id;
};

// const exportCSV = () => {
//   dt.value.exportCSV();
// };
const confirmDeleteSelected = () => {
  deleteMySwitchesDialog.value = true;
};
const deleteSelectedMySwitches = () => {
  myswitches.value = myswitches.value.filter(
    (val) => !selectedMySwitches.value.includes(val)
  );
  deleteMySwitchesDialog.value = false;
  selectedMySwitches.value = null;
  toast.add({
    severity: "success",
    summary: "Successful",
    detail: "MySwitches Deleted",
    life: 3000,
  });
};

// const getStatusLabel = (status) => {
//   switch (status) {
//     case "INSTOCK":
//       return "success";

//     case "LOWSTOCK":
//       return "warn";

//     case "OUTOFSTOCK":
//       return "danger";

//     default:
//       return null;
//   }
// };
</script>

<template>
    <div>
    <MegaMenu
      :model="menubar"
      class="menubar"
      
    >
      <template #start>
        <p class="backButton">
          <i class="pi pi-angle-left" style="font-size: 3rem" @click="navigateToPage"></i>  
        </p>
        <p class="title" style="">
          <h1 class="Settings">Settings</h1>
        </p>
      </template>
      <template #end>
      </template>
    </MegaMenu>
  </div>
  <div>
    <div class="card">
      <DataTable class="datatable"
        ref="dt"
        v-model:selection="selectedMySwitches"
        :value="myswitches"
        dataKey="IPaddress"
        :paginator="true"
        :rows="10"
        :filters="filters"
        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
        :rowsPerPageOptions="[5, 10, 25]"
        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} Switches"
      >
        <template #header>
          <div class="">
            <h4 class="tabletitle">Manage Switches</h4>
          </div>
          <Toolbar class="toolbar">
        <template #start>
            <IconField>
              <InputIcon>
                <i class="pi pi-search" />
              </InputIcon>
              <InputText
                v-model="filters['global'].value"
                placeholder="Search..."
              />
            </IconField>      
        </template>
        <template #end>
            <Button
            icon="pi pi-plus"
            outlined
            rounded
            class="roundButton"
            @click="openNew"
            style="margin-right: 10px;"
          />
            <Button
            icon="pi pi-trash"
            severity="danger"
            outlined
            rounded
            class ="roundButton"
            @click="confirmDeleteSelected"
            :disabled="!selectedMySwitches || !selectedMySwitches.length"
            style="margin-right: 0px;"
          />
            <!-- <FileUpload
            uploadIcon="pi pi-file-import"
            mode="basic"
            accept="image/*"
            :maxFileSize="1000000"
            class="custom-edit mr-2"
            auto
          />
            <Button
            icon="pi pi-file-export"
            severity="help"
            @click="exportCSV($event)"
            style="margin-left: 10px;"
            class="custom-edit mr-2"
          /> -->
        </template>
      </Toolbar>
        </template>
        <Column
          selectionMode="multiple"
          style="width: 3rem"
          :exportable="false"
        ></Column>
        <Column
          field="IPaddress"
          header="IP Address"
          sortable
          style="min-width: 12rem"
        ></Column>
        <Column
          field="name"
          header="Name"
          sortable
          style="min-width: 16rem"
        ></Column>
        <Column :exportable="false" style="min-width: 12rem">
          <template #body="slotProps">
             <div style="display: flex; justify-content: flex-end;">
              <Button
                icon="pi pi-pencil"
                outlined
                rounded
                class="roundButton"
                @click="editMySwitch(slotProps.data)"
                style="margin-right: 5px;"
              />
              <Button
                icon="pi pi-trash"
                outlined
                rounded
                severity="danger"
                class="roundButton"
                @click="confirmDeleteMySwitch(slotProps.data)"
                style="margin-left: 5px;"
              />
            </div>
          </template>
        </Column>
      </DataTable>
    </div>

    <Dialog
      v-model:visible="myswitchDialog"
      :style="{ width: '450px' }"
      header="Switch Details"
      :modal="true"
    >
      <div class="flex flex-col gap-6">
        
        <div class="input-name">
          <label for="name" class="block font-bold mb-3">Name</label>
          <InputText
            id="name"
            v-model.trim="myswitch.name"
            required="true"
            autofocus
            :invalid="submitted && !myswitch.name"
            fluid
          />
          <small v-if="submitted && !myswitch.name" class="text-red-500"
            >Name is required.</small
          >
        </div><div class="input-ip">
          <label for="IP Address" class="block font-bold mb-3">IP Address</label>
          <InputText
            id="IP Address"
            v-model.trim="myswitch.IPaddress"
            required="true"
            autofocus
            :invalid="submitted && !myswitch.IPaddress"
            fluid
          />
          <small v-if="submitted && !myswitch.name" class="text-red-500"
            >IP Address is required.</small
          >
        </div>    
      </div>
      
      <template #footer>
        <Button label="Cancel" icon="pi pi-times" text @click="hideDialog" />
        <Button label="Save" icon="pi pi-check" @click="saveMySwitch" />
      </template>
    </Dialog>

    <Dialog
      v-model:visible="deleteMySwitchDialog"
      :style="{ width: '450px' }"
      header="Confirm"
      :modal="true"
    >
      <div class="flex items-center gap-4">
        <span v-if="myswitch"
          >Are you sure you want to delete <b>{{ myswitch.name }}</b
          >?</span
        >
      </div>
      <template #footer>
        <Button
          label="No"
          icon="pi pi-times"
          text
          @click="deleteMySwitchDialog = false"
        />
        <Button label="Yes" icon="pi pi-check" @click="deleteMySwitch" />
      </template>
    </Dialog>

    <Dialog
      v-model:visible="deleteMySwitchesDialog"
      :style="{ width: '450px' }"
      header="Confirm"
      :modal="true"
    >
      <div class="flex items-center gap-4">
        <span v-if="myswitch"
          >Are you sure you want to delete the selected myswitches?</span
        >
      </div>
      <template #footer>
        <Button
          label="No"
          icon="pi pi-times"
          text
          @click="deleteMySwitchesDialog = false"
        />
        <Button
          label="Yes"
          icon="pi pi-check"
          text
          @click="deleteSelectedMySwitches"
        />
      </template>
    </Dialog>
  </div>
</template>

<style>
.input-name{
    margin-bottom: 15px;
}
.input-ip{
    margin-bottom: 15px;
}
.datatable{
    margin: 0px 86px;
}

.custom-edit{
    width: 40px;
    height: 40px;
    border-radius: 20px !important;
}

.custom-edit .p-button-label{
    display: none;
}
.toolbar{
    border: none !important;
    padding: 0px !important;
}

.pi-angle-left {
  padding-right: 0px; 
  cursor: pointer;
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.pi-angle-left:hover { 
  transform: translateY(-2px);
}
</style>