export const MySwitchService = {
  getMySwitchesData() {
    return [
      {
        id: "1000",
        name: "Bamboo Watch",
        inventoryStatus: "INSTOCK",
        IPaddress: "10.139.50.52",
      },
      {
        id: "1027",
        name: "Yellow Earbuds",
        inventoryStatus: "INSTOCK",
        IPaddress: "10.139.50.53",
      },
      {
        id: "1028",
        name: "Yoga Mat",
        inventoryStatus: "INSTOCK",
        IPaddress: "10.139.50.54",
      },
      {
        id: "1029",
        name: "Yoga Set",
        inventoryStatus: "INSTOCK",
        IPaddress: "10.139.50.55",
      },
    ];
  },

  getMySwitches() {
    return Promise.resolve(this.getMySwitchesData());
  }
  
};

export default MySwitchService;
