## How Its Works

MinIO is a high-performance, distributed object storage server designed specifically for large-scale private cloud infrastructures. It enables efficient storage management by aggregating persistent volumes (PVs) into scalable, distributed object storage. MinIO is fully compatible with Amazon S3 REST APIs, making it easy to integrate with existing applications and workflows.

**Steps to Install MinIO on Docker**

1. **Pull the MinIO Docker image**

Run the following command to download the MinIO Docker image:
```
docker pull minio/minio
```

2. **Create a directory for data storage:**

Choose a location on your host machine to store MinIO data. For example:
```
mkdir -p ~/minio/data
```

3. **Run the MinIO container:**

Use the docker run command to start the MinIO container. Replace `<your-access-key>` and `<your-secret-key>` with your desired credentials.
```
docker run -d --name minio \
  -p 9000:9000 \
  -p 9001:9001 \
  -v ~/minio/data:/data \
  -e "MINIO_ROOT_USER=<your-access-key>" \
  -e "MINIO_ROOT_PASSWORD=<your-secret-key>" \
  minio/minio server /data --console-address ":9001"
```

4. **Access the MinIO console:**

Once the container is running, you can access the MinIO web console by navigating to `http://localhost:9001` in your web browser. Log in using the access key and secret key you configured.