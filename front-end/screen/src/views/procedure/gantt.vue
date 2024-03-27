<template>
  <div id="gantt">
    <div class="bg-color-black">
      <!--      <div class="gantt-title">{{ this.title }}</div>-->
      <div :style="{ width: '100%', height: '100%' }" id="ganttChart">
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import Highcharts from "highcharts";
import exporting from "highcharts/modules/exporting";
import xrange from "highcharts/modules/xrange";

exporting(Highcharts);
xrange(Highcharts);

export default defineComponent({
  name: "gantt",
  data() {
    return {
      title: "gantt",
      categories: ['task1', 'task2', 'task3', 'task4'],
    }
  },
  components: {},
  mounted() {
    this.draw()
  },
  methods: {
    draw() {
      Highcharts.chart("ganttChart", {
        chart: {
          type: 'xrange',
          backgroundColor: 'rgba(0,0,0,0)',
        },
        title: {
          text: '工序甘特图',
          style: {
            color: '#008000FF',
            fontSize: '30px',
            fontWeight: 'bold'
          }
        },
        xAxis: {
          type: 'datetime',
          dateTimeLabelFormats: {
            week: '%Y/%m/%d'
          },
          labels: {
            style: {
              color: '#FFF',
              fontSize: '20px',
              fontWeight: 'bold'
            }
          }
        },
        yAxis: {
          title: {
            text: ''
          },
          labels: {
            style: {
              color: '#FFF',
              fontSize: '20px',
              fontWeight: 'bold'
            }
          },
          categories: this.categories,
          reversed: true
        },
        tooltip: {
          borderRadius: 10,
          borderWidth: 2,
          dateTimeLabelFormats: {
            day: '%Y/%m/%d'
          },
          style: {
            fontSize: '22px',
            fontWeight: 'bold'
          }
        },
        legend: {
          enabled: false
        },
        series: [{
          name: 'Task',
          // pointPadding: 0,
          // groupPadding: 0,
          pointWidth: 20,
          data: [{
            x: Date.UTC(2014, 10, 21),
            x2: Date.UTC(2014, 11, 2),
            y: 0
          }, {
            x: Date.UTC(2014, 11, 2),
            x2: Date.UTC(2014, 11, 5),
            y: 1
          }, {
            x: Date.UTC(2014, 11, 8),
            x2: Date.UTC(2014, 11, 9),
            y: 2
          }, {
            x: Date.UTC(2014, 11, 9),
            x2: Date.UTC(2014, 11, 19),
            y: 1
          }, {
            x: Date.UTC(2014, 11, 10),
            x2: Date.UTC(2014, 11, 23),
            y: 2
          }, {
            x: Date.UTC(2014, 11, 10),
            x2: Date.UTC(2014, 11, 13),
            y: 3
          },
            {
              x: Date.UTC(2014, 10, 24),
              x2: Date.UTC(2014, 11, 6),
              y: 3
            }
          ],
          dataLabels: {
            enabled: true
          }
        }]
      });

    }
  }
})
</script>

<style scoped lang="scss">
$box-height: 900px;
$box-width: 100%;
#gantt {
  padding: 20px 16px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 35px;
    border-radius: 10px;
  }
}
</style>
