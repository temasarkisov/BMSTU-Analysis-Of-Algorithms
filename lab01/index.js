"use strict";

const readlineSync = require('readline-sync');
const now = require("performance-now")

function removeLastSymbol(str) {
    return str.slice(0, -1);
}

function recursiveLevenshtein(strA, strB) {
    if (strA.length === 0 || strB.length === 0) {
        return Math.abs(strA.length - strB.length);
    }
    let cost = (strA.charAt(strA.length - 1) === strB.charAt(strB.length - 1)) ? 0 : 1;
    return Math.min(recursiveLevenshtein(strA, removeLastSymbol(strB)) + 1,
                    recursiveLevenshtein(removeLastSymbol(strA), strB) + 1,
                    recursiveLevenshtein(removeLastSymbol(strA), removeLastSymbol(strB)) + cost);
}

function matrixInitialization(str) {
    let matrix = [];
    for (let i = 0; i < str.length + 1; i++) {
        matrix[i] = [];
    }
    return matrix;
}

function recursiveOptimizedLevenshtein(strA, strB, matrix) {
    if (typeof(matrix) == 'undefined') {
        matrix = matrixInitialization(strA);
    }
    if (typeof(matrix[strA.length][strB.length]) !== 'undefined') {
        return matrix[strA.length][strB.length];
    } 
    if (strA.length === 0 || strB.length === 0) {
        matrix[strA.length][strB.length] = Math.abs(strA.length - strB.length);
        return matrix[strA.length][strB.length];
    }
    let cost = (strA.charAt(strA.length - 1) === strB.charAt(strB.length - 1)) ? 0 : 1;
    matrix[strA.length][strB.length] = Math.min(
        recursiveOptimizedLevenshtein(strA, removeLastSymbol(strB), matrix) + 1,
        recursiveOptimizedLevenshtein(removeLastSymbol(strA), strB, matrix) + 1,
        recursiveOptimizedLevenshtein(removeLastSymbol(strA), removeLastSymbol(strB), matrix) + cost);
    return matrix[strA.length][strB.length];
}

function matrixLevenshtein(strA, strB) {
    let matrix = [];
    for (let i = 0; i < strA.length + 1; i++) {
        matrix[i] = [];
        for (let j = 0; j < strB.length + 1; j++) {
            if (i === 0) {
                matrix[i][j] = j;
            } else if(j === 0 && i > 0) {
                matrix[i][j] = i;
            } else {
                let cost = (strA.charAt(i - 1) === strB.charAt(j - 1)) ? 0 : 1;
                matrix[i][j] = Math.min(matrix[i][j - 1] + 1,
                                        matrix[i - 1][j] + 1,
                                        matrix[i - 1][j - 1] + cost);
            }
        }
    }
    //console.log(matrix);
    return matrix[strA.length][strB.length];
}

function matrixDamerauLevenshtein(strA, strB) {
    let matrix = [];
    for (let i = 0; i < strA.length + 1; i++) {
        matrix[i] = [];
        for (let j = 0; j < strB.length + 1; j++) {
            if (i === 0 || j === 0) {
                matrix[i][j] = Math.max(i, j);
            } else {
                let cost = (strA.charAt(i - 1) === strB.charAt(j - 1)) ? 0 : 1;
                matrix[i][j] = Math.min(matrix[i][j - 1] + 1,
                    matrix[i - 1][j] + 1,
                    matrix[i - 1][j - 1] + cost);
                if (j > 1 && i > 1 && 
                    strA.charAt(i) === strB.charAt(j - 1) && 
                    strA.charAt(i - 1) === strB.charAt(j)) {
                        matrix[i][j] = Math.min(matrix[i][j], matrix[i - 2][j - 2] + 1);
                }
            }
        }
    }
    //console.log(matrix);
    return matrix[strA.length][strB.length];
}

//console.log(recursiveLevenshtein("hyundai", "honda"));
//console.log(matrixLevenshtein("hyundai", "honda"));
//console.log(matrixDamerauLevenshtein("hyundai", "honda"));
//console.log(recursiveOptimizedLevenshtein("hyundai", "honda"));

let menuOperation = -1;
let strA = "";
let strB = "";
let distance = 0;
let start = 0;
let end = 0;
while (menuOperation !== 0) {
    console.log("1 - Нахождение расстояния Левенштейна рекурсивно");
    console.log("2 - Нахождение расстояния Левенштейна рекурсивно с матрицей");
    console.log("3 - Нахождение расстояния Левенштейна матрично");
    console.log("4 - Нахождение расстояния Дамерау - Левенштейна матрично");
    console.log("0 - Выход");
    console.log('\n');
    menuOperation = readlineSync.question("Введите номер операции: ");
    menuOperation = parseInt(menuOperation);
    if (menuOperation !== 0 && (menuOperation === 1 || menuOperation === 2 || menuOperation === 3 || menuOperation === 4)) {
        strA = readlineSync.question("Введите первую строку: ");
        strB = readlineSync.question("Введите вторую строку: ");
        if (menuOperation === 1) {
            start = now();
            distance = recursiveLevenshtein(strA, strB);
            end = now();
        }
        else if (menuOperation === 2) {
            start = now();
            distance = recursiveOptimizedLevenshtein(strA, strB);
            end = now();
        }
        else if (menuOperation === 3) {
            start = now();
            distance = matrixLevenshtein(strA, strB);
            end = now();
        }
        else if (menuOperation === 4) {
            start = now();
            distance =  matrixDamerauLevenshtein(strA, strB);
            end = now();
        }
        console.log("\nРасстояние Левенштейна - " + distance);
        console.log("Время выполнения алгоритма - " + (end - start).toFixed(3) + " миллисекунд\n");
    }
    else if (menuOperation !== 0) {
        console.log("Неверный номер операции");
        menuOperation = 0;
    }
}