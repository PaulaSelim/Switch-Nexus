export const MySwitchService = {
  getMySwitchesData() {
    return [
      {
        id: "1000",
        code: "f230fh0g3",
        name: "Bamboo Watch",
        description: "MySwitch Description",
        image: "bamboo-watch.jpg",
        price: 65,
        category: "Accessories",
        quantity: 24,
        inventoryStatus: "INSTOCK",
        rating: 5,
      },
      {
        id: "1027",
        code: "acvx872gc",
        name: "Yellow Earbuds",
        description: "MySwitch Description",
        image: "yellow-earbuds.jpg",
        price: 89,
        category: "Electronics",
        quantity: 35,
        inventoryStatus: "INSTOCK",
        rating: 3,
      },
      {
        id: "1028",
        code: "tx125ck42",
        name: "Yoga Mat",
        description: "MySwitch Description",
        image: "yoga-mat.jpg",
        price: 20,
        category: "Fitness",
        quantity: 15,
        inventoryStatus: "INSTOCK",
        rating: 5,
      },
      {
        id: "1029",
        code: "gwuby345v",
        name: "Yoga Set",
        description: "MySwitch Description",
        image: "yoga-set.jpg",
        price: 20,
        category: "Fitness",
        quantity: 25,
        inventoryStatus: "INSTOCK",
        rating: 8,
      },
    ];
  },

  getMySwitchesWithOrdersData() {
    return [
      {
        id: "1000",
        code: "f230fh0g3",
        name: "Bamboo Watch",
        description: "MySwitch Description",
        image: "bamboo-watch.jpg",
        price: 65,
        category: "Accessories",
        quantity: 24,
        inventoryStatus: "INSTOCK",
        rating: 5,
        orders: [
          {
            id: "1000-0",
            myswitchCode: "f230fh0g3",
            date: "2020-09-13",
            amount: 65,
            quantity: 1,
            customer: "David James",
            status: "PENDING",
          },
          {
            id: "1000-1",
            myswitchCode: "f230fh0g3",
            date: "2020-05-14",
            amount: 130,
            quantity: 2,
            customer: "Leon Rodrigues",
            status: "DELIVERED",
          },
          {
            id: "1000-2",
            myswitchCode: "f230fh0g3",
            date: "2019-01-04",
            amount: 65,
            quantity: 1,
            customer: "Juan Alejandro",
            status: "RETURNED",
          },
          {
            id: "1000-3",
            myswitchCode: "f230fh0g3",
            date: "2020-09-13",
            amount: 195,
            quantity: 3,
            customer: "Claire Morrow",
            status: "CANCELLED",
          },
        ],
      },
      {
        id: "1027",
        code: "acvx872gc",
        name: "Yellow Earbuds",
        description: "MySwitch Description",
        image: "yellow-earbuds.jpg",
        price: 89,
        category: "Electronics",
        quantity: 35,
        inventoryStatus: "INSTOCK",
        rating: 3,
        orders: [
          {
            id: "1027-0",
            myswitchCode: "acvx872gc",
            date: "2020-01-29",
            amount: 89,
            quantity: 1,
            customer: "Veronika Inouye",
            status: "DELIVERED",
          },
          {
            id: "1027-1",
            myswitchCode: "acvx872gc",
            date: "2020-06-11",
            amount: 89,
            quantity: 1,
            customer: "Willard Kolmetz",
            status: "DELIVERED",
          },
        ],
      },
    ];
  },

  getMySwitchesMini() {
    return Promise.resolve(this.getMySwitchesData().slice(0, 5));
  },

  getMySwitchesSmall() {
    return Promise.resolve(this.getMySwitchesData().slice(0, 10));
  },

  getMySwitches() {
    return Promise.resolve(this.getMySwitchesData());
  },

  getMySwitchesWithOrdersSmall() {
    return Promise.resolve(this.getMySwitchesWithOrdersData().slice(0, 10));
  },

  getMySwitchesWithOrders() {
    return Promise.resolve(this.getMySwitchesWithOrdersData());
  },
};

export default MySwitchService;
