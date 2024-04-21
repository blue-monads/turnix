export const formatCurrency = (price: string | number) => {

    const numberPrice = typeof price === "number" ? Number(price) : parseFloat(price);
    if (isNaN(numberPrice)) {
        return price;
    }

    return numberPrice.toLocaleString(undefined, {
        minimumFractionDigits: 2,
        maximumFractionDigits: 2
    });
}
