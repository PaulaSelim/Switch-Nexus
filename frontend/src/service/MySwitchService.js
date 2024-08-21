export const MySwitchService = {
  getMySwitchesData() {
    const switches = JSON.parse(localStorage.getItem("myswitches") || "[]");
    return switches.length > 0 ? switches : [
    ];
  },

  getMySwitches() {
    return Promise.resolve(this.getMySwitchesData());
  },
  
  saveMySwitches(switches) {
    localStorage.setItem("myswitches", JSON.stringify(switches));
  }
};

export default MySwitchService;
