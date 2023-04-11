// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file generates the area of a triangle
"use strict"
function calculate() {
  // input
  const radius = parseInt(document.getElementById("radius").value)
  const round = 4 / 3

  // process
  const volume = round * Math.PI * radius ** 3

  // output
  document.getElementById("volume").innerHTML = "Volume is: " + volume.toFixed(2) + " cm³"
}
