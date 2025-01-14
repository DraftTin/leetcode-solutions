type TimeMap struct {
  Mp map[int]map[string]string 
  TimeLine []int
}


func Constructor() TimeMap {
  timeMap := TimeMap{}
  timeMap.Mp = map[int]map[string]string{}
  timeMap.TimeLine = []int{}
  return timeMap 
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
  if _, ok := this.Mp[timestamp]; ok == false {
    this.Mp[timestamp] = map[string]string{}
  } 
  this.Mp[timestamp][key] = value
  this.TimeLine = append(this.TimeLine, timestamp)

}


func (this *TimeMap) Get(key string, timestamp int) string {
  index := sort.Search(len(this.TimeLine), func(i int) bool {
    return this.TimeLine[i] >= timestamp
  })
  if index < len(this.TimeLine) && this.TimeLine[index] == timestamp  {
      return this.Mp[timestamp][key]
  }
  for i := index - 1; i >= 0; i-- {
    if val, ok := this.Mp[this.TimeLine[i]][key]; ok == true {
      return val
    } 
  }
  return ""
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
