import http from 'k6/http';

export const options = {
  duration: '60s',
  vus: 10,
};

export default function () {
  http.get(
    'http://localhost:8081/dorm?capacity=2&lower_price=2500&upper_price=7000&dorm_facilities=Laundry&dorm_facilities=Parking%20Lot&room_facilities=A%2FC&room_facilities=TV&room_facilities=Refrigerator'
  );
}
