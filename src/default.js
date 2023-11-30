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
    document.getElementById("prev-btn").innerHTML = "submit";
    document.getElementById("next-btn").innerHTML = "su";
  } else {
    document.getElementById("prev-btn").innerHTML = "Next";
  }

  fixStepIndicator(n);
}

function nextPrev(n) {
  var x = document.getElementsByClassName("tab");
  if (n == 1 && !validateForm()) return false;

  if (x[currentTab] == 2) {
    x[currentTab].id = "tab-2";
  }

  if (x[currentTab] == 3) {
    x[currentTab].id = "tab-3";
  }

  x[currentTab].style.display = "none";
  currentTab = currentTab + n;

  if (currentTab >= x.length) {
    document.getElementById("form").submit();
    return false;
  }

  showTab(currentTab);
}

function validateForm() {
  var x,
    y,
    i,
    valid = true;

  x = document.getElementsByClassName("tab");
  y = x[currentTab].getElementsByTagName("input");

  for (i = 0; i < y.length; i++) {
    if (y[i].value == "") {
      y[i].className += " invalid";
      valid = false;
    }
  }

  if (valid) {
    document.getElementsByClassName("step")[currentTab].className += " finish";
  }
  return valid;
}

function fixStepIndicator(n) {
  var i,
    x = document.getElementsByClassName("step");
  for (i = 0; i < x.length; i++) {
    x[i].className = x[i].className.replace(" active", "");
  }

  x[n].className += " active";
}
