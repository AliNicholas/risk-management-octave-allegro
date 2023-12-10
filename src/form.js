var currentTab = 0;
showTab(currentTab);

function showTab(n) {
  var x = document.getElementsByClassName("tab");

  x[n].style.display = "block";

  if (n == 0) {
    document.getElementById("prev-btn").style.display = "none";
  } else {
    document.getElementById("prev-btn").style.display = "inline";
  }

  if (n == x.length - 1) {
    // document.getElementById("next-btn").innerHTML = "Submit";
    document.getElementById("next-btn").style.display = "none";
    document.getElementById("submit-btn").style.display = "inline";
  } else {
    document.getElementById("next-btn").style.display = "inline";
    document.getElementById("submit-btn").style.display = "none";
  }

  fixStepIndicator(n);
}

function changeTab(n) {
  var x = document.getElementsByClassName("tab");
  if (n == 1 && !validateForm()) return false;

  x[currentTab].id = "";

  x[currentTab].style.display = "none";
  currentTab = currentTab + n;

  x[currentTab].id = "tab-" + (currentTab + 1);
  showTab(currentTab);
}

function validateForm() {
  var x,
    y,
    i,
    valid = true;

  x = document.getElementsByClassName("tab");
  y = x[currentTab].querySelectorAll("input, select, textarea, option");

  for (i = 0; i < y.length; i++) {
    if (y[i].value.trim() === "") {
      y[i].classList.add("invalid");
      valid = false;
    }
  }

  if (valid) {
    document.getElementsByClassName("step")[currentTab].classList.add("finish");
  }
  return valid;
}

function fixStepIndicator(n) {
  var i,
    x = document.getElementsByClassName("step");
  for (i = 0; i < x.length; i++) {
    x[i].classList.remove("active");
  }

  x[n].classList.add("active");
}

// document.addEventListener("DOMContentLoaded", function () {
//   // Daftar semua elemen select
//   var selects = document.querySelectorAll("select");

//   // Tambahkan event listener untuk setiap elemen select
//   selects.forEach(function (select) {
//     select.addEventListener("change", function () {
//       // Ambil nilai yang dipilih
//       var selectedValue = this.value;

//       // Nonaktifkan nilai yang dipilih dalam dropdown lainnya
//       selects.forEach(function (otherSelect) {
//         if (otherSelect !== select) {
//           // Aktifkan kembali semua opsi
//           otherSelect.querySelectorAll("option").forEach(function (option) {
//             option.disabled = false;
//           });

//           // Nonaktifkan opsi yang sudah dipilih di dropdown lainnya
//           if (selectedValue !== "None") {
//             otherSelect.querySelector(
//               "option[value='" + selectedValue + "']"
//             ).disabled = true;
//           }
//         }
//       });
//     });
//   });
// });
