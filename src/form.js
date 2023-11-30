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

// select options
// Fungsi untuk mengubah nilai opsi pada kategori lain
function updateOptions(selectedCategory, selectedValue) {
  // Ambil semua elemen select
  var selects = document.getElementsByTagName("select");

  // Iterasi melalui elemen select
  for (var i = 0; i < selects.length; i++) {
    var currentCategory = selects[i].id;

    // Jika bukan kategori yang dipilih, perbarui nilai opsi
    if (currentCategory !== selectedCategory) {
      // Dapatkan semua opsi dalam kategori tersebut
      var options = selects[i].getElementsByTagName("option");

      // Iterasi melalui opsi dan nonaktifkan opsi yang dipilih di kategori lain
      for (var j = 0; j < options.length; j++) {
        if (options[j].value === selectedValue) {
          options[j].disabled = true;
        }
      }
    }
  }
}

// Fungsi untuk menangani perubahan pada opsi
function handleOptionChange(event) {
  var selectedCategory = event.target.id;
  var selectedValue = event.target.value;

  // Reset semua opsi
  resetOptions();

  // Perbarui opsi pada kategori lain
  updateOptions(selectedCategory, selectedValue);
}

// Fungsi untuk mereset nilai opsi
function resetOptions() {
  // Ambil semua elemen select
  var selects = document.getElementsByTagName("select");

  // Iterasi melalui elemen select
  for (var i = 0; i < selects.length; i++) {
    // Dapatkan semua opsi dalam kategori tersebut
    var options = selects[i].getElementsByTagName("option");

    // Iterasi melalui opsi dan aktifkan kembali semua opsi
    for (var j = 0; j < options.length; j++) {
      options[j].disabled = false;
    }
  }
}

// Tambahkan event listener pada setiap elemen select
var selects = document.getElementsByTagName("select");
for (var i = 0; i < selects.length; i++) {
  selects[i].addEventListener("change", handleOptionChange);
}
