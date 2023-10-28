export default function formatDate(dataISO: string) {
  // Cria um objeto Date a partir da data no formato ISO
  const data = new Date(dataISO);

  // Extrai o dia, mês e ano da data
  const dia = data.getDate();
  const mes = data.getMonth() + 1; // Lembre-se que os meses em JavaScript começam em 0
  const ano = data.getFullYear();

  // Formata o dia e o mês para ter dois dígitos
  const diaFormatado = dia.toString().padStart(2, "0");
  const mesFormatado = mes.toString().padStart(2, "0");

  // Retorna a data formatada como DD/MM/YYYY
  return `${diaFormatado}/${mesFormatado}/${ano}`;
}
