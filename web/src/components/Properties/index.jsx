import React from "react";
import { useQuery } from "react-query";
import { Card, Row, Col, Input, Image, Carousel } from "antd";
import { RiseOutlined, FallOutlined } from "@ant-design/icons";

const Properties = () => {
  const { isLoading, error, data } = useQuery("propertyData", () =>
    fetch("http://localhost:8080/properties").then((res) => res.json())
  );
  if (isLoading) return "Loading...";

  if (error) return "An error has occurred: " + error.message;

  console.log(data);
  return (
    <div>
      <Row gutter={[32, 32]} className="crypto-card-container">
        {data?.map((property) => (
          <Col
            xs={24}
            sm={12}
            lg={6}
            className="property-card"
            key={property.listingid}
          >
            {/* <Link to={`/crypto/${currency.coin_id}/${currency.coin_uuid}`}> */}
            <Card
              hoverable="true"
              // style={millify(currency.change) < 0 ? { backgroundColor: 'red' } : { backgroundColor: 'green' }}
            >
              <Image.PreviewGroup>
                <Carousel autoplay>
                  {property.images.map((image, index) => {
                    return (
                      <Image
                        key={index}
                        src={image}
                        preview={{ getContainer: "#root" }}
                      />
                    );
                  })}
                </Carousel>
              </Image.PreviewGroup>
              <div className="property-details">
                {/* <a>
                  <href></href>
                </a> */}
                <p>
                  <b>{property.title}</b>
                </p>
                <a
                  href={property.listingurl}
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Link
                </a>

                <p>Price: {property.price}</p>
                <p>Published On: {property.publishedon}</p>

                {property.availablefrom !== "" ? (
                  <p>Available: {property.availablefrom}</p>
                ) : (
                  <></>
                )}
                {property.bedrooms !== "" ? (
                  <p>Bedrooms: {property.bedrooms}</p>
                ) : (
                  <></>
                )}
                {property.bathrooms !== "" ? (
                  <p>Bathrooms: {property.bathrooms}</p>
                ) : (
                  <></>
                )}
              </div>
            </Card>
            {/* </Link> */}
          </Col>
        ))}
      </Row>
    </div>
  );
};

export default Properties;
